#!/usr/bin/env python3

from pathlib import Path

from src.app import App

if __name__ == "__main__":
    glade = (Path("glade") / "app.glade").as_posix()
    App(glade).run()
