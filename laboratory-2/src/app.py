import xml.etree.ElementTree as ET
from dataclasses import dataclass
from pathlib import Path
from time import gmtime, strftime
from typing import Callable

import src.text_placeholders
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

        # XML file parameters.
        self._xml_name = "state.xml"

        # Init buffers.
        self._buffer_maze = self._builder.get_object("buffer_maze")
        self._buffer_text = self._builder.get_object("buffer_text")

        # Init maze.
        self._maze = Maze()
        self._redraw_maze()

        self._moves = {
            "U": Move(
                move=self._maze.U,
                text_succ=src.text_placeholders.SUCC_U,
                text_fail=src.text_placeholders.FAIL_U,
            ),
            "D": Move(
                move=self._maze.D,
                text_succ=src.text_placeholders.SUCC_D,
                text_fail=src.text_placeholders.FAIL_D,
            ),
            "L": Move(
                move=self._maze.L,
                text_succ=src.text_placeholders.SUCC_L,
                text_fail=src.text_placeholders.FAIL_L,
            ),
            "R": Move(
                move=self._maze.R,
                text_succ=src.text_placeholders.SUCC_R,
                text_fail=src.text_placeholders.FAIL_R,
            ),
        }

        # Connect handlers.
        self._builder.connect_signals(
            {
                "onDestroy": Gtk.main_quit,
                "onButtonU": self._onButtonU,
                "onButtonD": self._onButtonD,
                "onButtonL": self._onButtonL,
                "onButtonR": self._onButtonR,
            }
        )

    def run(self):
        self._builder.get_object("main_window").show_all()
        Gtk.main()

    def _append_into_text_window(self, text):
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

    def _xml_load(self, xml_file=None):
        pass

    def _xml_dump(self):
        # Parameters.
        x, y = self._maze.get_position()

        root = ET.Element("root")

        # Dump player info.
        player = ET.SubElement(root, "player")
        ET.SubElement(player, "x").text = str(x)
        ET.SubElement(player, "y").text = str(y)

        # Dump current maze.
        maze = ET.SubElement(root, "maze")
        ET.SubElement(maze, "map").text = str(x)

        # Dump file.
        tree = ET.ElementTree(root)
        tree.write(Path("xml") / self._xml_name)

    def _xml_delete(self):
        pass

    @staticmethod
    def _text_and_time(text):
        return strftime("%H:%M:%S", gmtime()) + " " + text

    def _redraw_maze(self):
        self._buffer_maze.set_text(self._maze.draw_text())

    def _move(self, move_object: Move) -> None:
        if move_object.move():
            self._redraw_maze()

            text = self._text_and_time(move_object.text_succ)
            self._append_into_text_window(text)

            self._xml_dump()

            if self._maze.is_hero_in_final():
                text = self._text_and_time(src.text_placeholders.TEXT_WIN)

                self._append_into_text_window(text)
                self._xml_delete()

            return

        text = self._text_and_time(move_object.text_fail)
        self._append_into_text_window(text)

    def _onButtonU(self, button):
        self._move(self._moves["U"])

    def _onButtonD(self, button):
        self._move(self._moves["D"])

    def _onButtonL(self, button):
        self._move(self._moves["L"])

    def _onButtonR(self, button):
        self._move(self._moves["R"])
