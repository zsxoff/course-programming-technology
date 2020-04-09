import numpy as np

from src.maze_placeholder import MAZE_MAP_CLASSIC, E, F, P, W


class Maze:
    def __init__(self):
        self.reset()

    def reset(self):
        self._maze_arr = MAZE_MAP_CLASSIC.copy()
        self._maze_w = self._maze_arr.shape[0]
        self._maze_h = self._maze_arr.shape[1]
        self._player_x = 1
        self._player_y = 1

        # Set final exit coords.
        wins_positions = np.argwhere(self._maze_arr == F)

        if len(wins_positions) > 1:
            raise Exception("Incorrect maze map")

        self._win_pos = wins_positions[0]

    def maze_tostring(self):
        return np.array2string(self._maze_arr, separator=",")

    # -------------------------------------------------------------------------

    def get_position(self):
        """
        Return player position.

        Returns:
            (tuple): Pair of coordinates in array.

        """
        return self._player_x, self._player_y

    def set_position(self, x, y):
        self._player_x, self._player_y = x, y

    def get_maze(self):
        return self._maze_arr

    def set_maze(self, maze_arr):
        self._maze_arr = maze_arr

    # -------------------------------------------------------------------------

    def _move(self, x, y):
        """
        Moves a player in a maze.

        Args:
            x (int): Coord X.
            y (int): Coord Y.

        Returns:
            bool: State whether the character was able to go.

        """
        # Exception for out of bound.
        if x < 0 or x > self._maze_h + 1 or y < 0 or y > self._maze_w + 1:
            return False

        # If empty space or final point ahead - go.
        if self._maze_arr[x][y] != W:
            return True

        # else, wall ahead.
        return False

    def is_hero_in_final(self):
        return (self._player_x, self._player_y) == (self._win_pos[0], self._win_pos[1],)

    def U(self):
        if self._move(self._player_x - 1, self._player_y):
            self._maze_arr[self._player_x][self._player_y] = E
            self._player_x -= 1
            self._maze_arr[self._player_x][self._player_y] = P
            return True
        return False

    def D(self):
        if self._move(self._player_x + 1, self._player_y):
            self._maze_arr[self._player_x][self._player_y] = E
            self._player_x += 1
            self._maze_arr[self._player_x][self._player_y] = P
            return True
        return False

    def L(self):
        if self._move(self._player_x, self._player_y - 1):
            self._maze_arr[self._player_x][self._player_y] = E
            self._player_y -= 1
            self._maze_arr[self._player_x][self._player_y] = P
            return True
        return False

    def R(self):
        if self._move(self._player_x, self._player_y + 1):
            self._maze_arr[self._player_x][self._player_y] = E
            self._player_y += 1
            self._maze_arr[self._player_x][self._player_y] = P
            return True
        return False

    # -------------------------------------------------------------------------

    def draw_text(self):
        # Text header.
        text_head = "Добро пожаловать в лабиринт!"

        # Convert NumPy array to text.
        text_maze = "\n".join(
            [" ".join([str(x) for x in row]) for row in self._maze_arr]
        )

        # Add help text.
        text_foot = "\n".join([f"{P} - персонаж", f"{W} - стена", f"{F} - выход"])

        # Generate full text.
        text_full = "\n\n".join([text_head, text_maze, text_foot])

        # Replace symbols in final text.
        replaces = [(str(E), " "), (str(W), "#"), (str(F), "F"), (str(P), "+")]

        for pair in replaces:
            text_full = text_full.replace(pair[0], pair[1])

        return text_full
