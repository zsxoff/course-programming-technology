from dataclasses import dataclass
from typing import Callable

from src.maze import Maze

import gi  # isort:skip

gi.require_version("Gtk", "3.0")  # isort:skip
from gi.repository import Gtk  # isort:skip


@dataclass
class Move:
    move: Callable
    succ: str
    fail: str


class App:
    def __init__(self, glade_file):
        self.builder = Gtk.Builder()
        self.builder.add_from_file(glade_file)

        # Init buffers.
        self._buffer_maze = self.builder.get_object("buffer_maze")
        self._buffer_text = self.builder.get_object("buffer_text")

        # Init maze.
        self._maze = Maze()
        self._redraw_maze()

        self.moves = {
            "U": Move(
                move=self._maze.U,
                succ="Вы сдвинулись вверх",
                fail="Сверху стена, вы остаетесь на месте",
            ),
            "D": Move(
                move=self._maze.D,
                succ="Вы сдвинулись вниз",
                fail="Снизу стена, вы остаетесь на месте",
            ),
            "L": Move(
                move=self._maze.L,
                succ="Вы сдвинулись влево",
                fail="Слева стена, вы остаетесь на месте",
            ),
            "R": Move(
                move=self._maze.R,
                succ="Вы сдвинулись вправо",
                fail="Справа стена, вы остаетесь на месте",
            ),
        }

        # Connect handlers.
        self.builder.connect_signals(
            {
                "onDestroy": Gtk.main_quit,
                "onButtonU": self._onButtonU,
                "onButtonD": self._onButtonD,
                "onButtonL": self._onButtonL,
                "onButtonR": self._onButtonR,
            }
        )

    def run(self):
        self.builder.get_object("main_window").show_all()
        Gtk.main()

    def _append_into_text_window(self, text):
        # Insert text into end of buffer.
        self._buffer_text.insert(self._buffer_text.get_end_iter(), text + "\n")

        # Scroll buffer at the end of text.
        self.builder.get_object("gtk_log").scroll_to_mark(
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

    def _redraw_maze(self):
        self._buffer_maze.set_text(self._maze.to_text())

    def _dump_to_xml(self):
        pass

    def _move(self, move_object: Move) -> None:
        if move_object.move():
            self._redraw_maze()
            self._append_into_text_window(move_object.succ)
        else:
            self._append_into_text_window(move_object.fail)

    def _onButtonU(self, button):
        self._move(self.moves["U"])

    def _onButtonD(self, button):
        self._move(self.moves["D"])

    def _onButtonL(self, button):
        self._move(self.moves["L"])

    def _onButtonR(self, button):
        self._move(self.moves["R"])
