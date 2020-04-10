import numpy as np

from src.maze_placeholder import MAZE_MAP_CLASSIC, E, F, P, W


class Maze:
    def __init__(self):
        self.reset()

    def reset(self):
        """Reset maze state."""
        self._maze_arr = MAZE_MAP_CLASSIC.copy()
        self._player_x = 1
        self._player_y = 1

        # Set final exit coords.
        wins_positions = np.argwhere(self._maze_arr == F)

        if len(wins_positions) > 1:
            raise Exception("Incorrect maze map")

        self._win_pos = wins_positions[0]

    def maze_to_string(self):
        """
        Dump flatten maze map as string vector.

        Returns:
            str: string vector of flatten map.

        """
        return ",".join([str(x) for x in self._maze_arr.flatten()])

    def maze_from_vector(self, vector, w, h):
        """
        Create 2D maze map from flatten vector.

        Args:
            vector (numpy.ndarray): Flatten map vector.
            w (int): Width of matrix.
            h (int): Height of matrix.

        """
        self._maze_arr = vector.reshape((w, h))

    def get_player_position(self):
        """
        Get player position.

        Returns:
            tuple: Pair of coordinates in array.

        """
        return self._player_x, self._player_y

    def set_player_position(self, x, y):
        """
        Set player position.

        Args:
            x (int): Position in row.
            y (int): Position in col.

        """
        self._player_x, self._player_y = x, y

    def get_maze_size(self):
        """
        Get maze size as tuple.

        Returns:
            tuple: Maze matrix shape as tuple of (rows, cols).

        """
        return self._maze_arr.shape

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
        maze_w = self._maze_arr.shape[0]
        maze_h = self._maze_arr.shape[1]

        if x < 0 or x > maze_h + 1 or y < 0 or y > maze_w + 1:
            return False

        # If empty space or final point ahead - go.
        if self._maze_arr[x][y] != W:
            return True

        # else, wall ahead.
        return False

    def is_player_in_final(self):
        """
        Check if player in final position.

        Returns:
            bool: True if player in final position else False.

        """
        return (self._player_x, self._player_y) == (
            self._win_pos[0],
            self._win_pos[1],
        )

    def move_u(self):
        """Move player up."""
        if self._move(self._player_x - 1, self._player_y):
            self._maze_arr[self._player_x][self._player_y] = E

            self._player_x -= 1
            self._maze_arr[self._player_x][self._player_y] = P

            return True

        return False

    def move_d(self):
        """Move player down."""
        if self._move(self._player_x + 1, self._player_y):
            self._maze_arr[self._player_x][self._player_y] = E

            self._player_x += 1
            self._maze_arr[self._player_x][self._player_y] = P

            return True

        return False

    def move_l(self):
        """Move player left."""
        if self._move(self._player_x, self._player_y - 1):
            self._maze_arr[self._player_x][self._player_y] = E

            self._player_y -= 1
            self._maze_arr[self._player_x][self._player_y] = P

            return True

        return False

    def move_r(self):
        """Move player right."""
        if self._move(self._player_x, self._player_y + 1):
            self._maze_arr[self._player_x][self._player_y] = E

            self._player_y += 1
            self._maze_arr[self._player_x][self._player_y] = P

            return True

        return False

    def draw_text(self):
        # Text header.
        text_head = "Добро пожаловать в лабиринт!"

        # Convert NumPy array to text.
        text_maze = "\n".join(
            [" ".join([str(x) for x in row]) for row in self._maze_arr]
        )

        # Add help text.
        text_foot = "\n".join([f"{P} - персонаж", f"{W} - стена"])

        # Generate full text.
        text_full = "\n\n".join([text_head, text_maze, text_foot])

        # Replace symbols in final text.
        replaces = [(str(E), " "), (str(W), "#"), (str(F), " "), (str(P), "*")]

        for pair in replaces:
            text_full = text_full.replace(pair[0], pair[1])

        return text_full
