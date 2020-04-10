import os
import xml.etree.ElementTree as ET
from dataclasses import dataclass
from pathlib import Path
from time import gmtime, strftime
from typing import Callable

import numpy as np

import src.placeholders_text
from src.maze import Maze

import gi  # isort:skip


gi.require_version("Gtk", "3.0")  # isort:skip
from gi.repository import Gtk  # isort:skip


@dataclass
class Move:
    move: Callable
    text_succ: str
    text_fail: str


class App:
    def __init__(self, glade_file):
        self._builder = Gtk.Builder()
        self._builder.add_from_file(glade_file)

        # Init buffers.
        self._buffer_maze = self._builder.get_object("buffer_maze")
        self._buffer_text = self._builder.get_object("buffer_text")

        # Init maze.
        self._maze = Maze()

        self._moves = {
            "U": Move(
                move=self._maze.move_u,
                text_succ=src.placeholders_text.SUCC_U,
                text_fail=src.placeholders_text.FAIL_U,
            ),
            "D": Move(
                move=self._maze.move_d,
                text_succ=src.placeholders_text.SUCC_D,
                text_fail=src.placeholders_text.FAIL_D,
            ),
            "L": Move(
                move=self._maze.move_l,
                text_succ=src.placeholders_text.SUCC_L,
                text_fail=src.placeholders_text.FAIL_L,
            ),
            "R": Move(
                move=self._maze.move_r,
                text_succ=src.placeholders_text.SUCC_R,
                text_fail=src.placeholders_text.FAIL_R,
            ),
        }

        # Connect handlers.
        self._builder.connect_signals(
            {
                "onDestroy": Gtk.main_quit,
                "onButtonU": self._button_move_u,
                "onButtonD": self._button_move_d,
                "onButtonL": self._button_move_l,
                "onButtonR": self._button_move_r,
                "onButtonReset": self._full_reset_callback,
            }
        )

        # XML file parameters.
        self._xml_name = "state.xml"
        self._xml_path = Path("xml") / self._xml_name
        self._xml_load_state()
        self._gui_replot_maze()

    def run(self):
        """Run GTK window."""
        self._builder.get_object("main_window").show_all()
        Gtk.main()

    def _gui_replot_maze(self):
        """Redraw current maze map in text buffer."""
        self._buffer_maze.set_text(self._maze.draw_text())

    def _gui_append_into_log(self, text):
        """
        Append text into log window, then scroll at the end of window.

        Args:
            text (str): Text for append.

        """
        # Insert text into end of buffer.
        self._buffer_text.insert(self._buffer_text.get_end_iter(), text + "\n")

        # Scroll buffer at the end of text.
        self._builder.get_object("gtk_log").scroll_to_mark(
            mark=self._buffer_text.create_mark(
                mark_name="",
                where=self._buffer_text.get_end_iter(),
                left_gravity=False,
            ),
            within_margin=0,
            use_align=False,
            xalign=0,
            yalign=0,
        )

    def _xml_dump(self):
        """
        Dump current state to XML file.

        Returns:
            bool: True if state successfully dumps else False.

        """
        root = ET.Element("root")

        # Parameters.
        x, y = self._maze.get_player_position()
        maze_txt = self._maze.maze_to_string()
        w, h = self._maze.get_maze_size()

        # Dump player info.
        player = ET.SubElement(root, "player")
        ET.SubElement(player, "x").text = str(x)
        ET.SubElement(player, "y").text = str(y)

        # Dump current maze.
        maze = ET.SubElement(root, "maze")
        ET.SubElement(maze, "map").text = maze_txt
        ET.SubElement(maze, "w").text = str(w)
        ET.SubElement(maze, "h").text = str(h)

        # Dump file.
        tree = ET.ElementTree(root)
        tree.write(self._xml_path)

        return True

    def _xml_load_state(self):
        """
        Load current XML state from file.

        Returns:
            bool: True if state successfully load else False.

        """
        if not os.path.exists(self._xml_path):
            return False

        tree = ET.parse(self._xml_path)
        root = tree.getroot()

        # Get player info.
        player = root.findall("player")[0]
        player_x = int(player.find("x").text)
        player_y = int(player.find("y").text)

        # Get maze.
        maze = root.findall("maze")[0]
        maze_w = int(maze.find("w").text)
        maze_h = int(maze.find("h").text)
        maze_map = maze.find("map").text

        # Init maze.
        maze_vector = np.fromstring(maze_map, sep=",", dtype=np.uint8)
        self._maze.maze_from_vector(maze_vector, maze_w, maze_h)
        self._maze.set_player_position(player_x, player_y)

        self._gui_replot_maze()

        # Say about state.
        text = self._text_and_time(src.placeholders_text.TEXT_XML)
        self._gui_append_into_log(text)

        return True

    def _xml_delete(self):
        """
        Delete current XML state file.

        Returns:
            bool: True if file removed else False.

        """
        if os.path.exists(self._xml_path):
            os.remove(self._xml_path)
            return True
        return False

    def _xml_rewrite(self):
        """
        Delete current XML state file and dump new.

        Returns:
            bool: True if file rewrited else False.

        """
        return all([self._xml_delete(), self._xml_dump()])

    @staticmethod
    def _text_and_time(text):
        return strftime("%H:%M:%S", gmtime()) + " " + text

    def _move_player(self, move_object):
        """
        Move player by state in move_object.

        Args:
            move_object (Move): move object with moving state.

        """
        if move_object.move():
            # If player get a move, redraw maze and dump XML state.

            self._gui_replot_maze()
            text = self._text_and_time(move_object.text_succ)
            self._gui_append_into_log(text)
            self._xml_dump()

            # If player in final position, drop XML state.

            if self._maze.is_player_in_final():
                text = self._text_and_time(src.placeholders_text.TEXT_WIN)
                self._gui_append_into_log(text)
                self._xml_delete()
        else:
            # Print fail message if wall ahead.

            text = self._text_and_time(move_object.text_fail)
            self._gui_append_into_log(text)

    # -------------------------------------------------------------------------

    def _button_move_u(self, button):
        _ = button
        self._move_player(self._moves["U"])

    def _button_move_d(self, button):
        _ = button
        self._move_player(self._moves["D"])

    def _button_move_l(self, button):
        _ = button
        self._move_player(self._moves["L"])

    def _button_move_r(self, button):
        _ = button
        self._move_player(self._moves["R"])

    def _full_reset_callback(self, button):
        _ = button
        self._maze.reset()
        self._xml_rewrite()
        self._gui_replot_maze()

        text = self._text_and_time(src.placeholders_text.TEXT_RELOAD)
        self._gui_append_into_log(text)
