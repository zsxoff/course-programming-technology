import numpy as np

E = 0
W = 1
F = 2
P = 3


class Maze:
    def __init__(self):
        self.maze = np.asarray(
            [
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, P, E, E, E, E, E, E, E, E, E, E, E, E, E, F, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
                [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
            ]
        )

        self.x = 1
        self.y = 1

    def move(self, x, y):
        if self.maze[x][y] != W:
            return True
        return False

    def U(self):
        if self.move(self.x - 1, self.y):
            self.maze[self.x][self.y] = E
            self.x -= 1
            self.maze[self.x][self.y] = P
            return True
        return False

    def D(self):
        if self.move(self.x + 1, self.y):
            self.maze[self.x][self.y] = E
            self.x += 1
            self.maze[self.x][self.y] = P
            return True
        return False

    def L(self):
        if self.move(self.x, self.y - 1):
            self.maze[self.x][self.y] = E
            self.y -= 1
            self.maze[self.x][self.y] = P
            return True
        return False

    def R(self):
        if self.move(self.x, self.y + 1):
            self.maze[self.x][self.y] = E
            self.y += 1
            self.maze[self.x][self.y] = P
            return True
        return False

    def _maze_encoder(self, text):
        replaces = [(str(E), " "), (str(W), "#"), (str(F), "F"), (str(P), "+")]

        for pair in replaces:
            text = text.replace(pair[0], pair[1])

        return text

    def to_text(self):
        text_head = "Добро пожаловать в олдскульный лабиринт!"

        text_maze = "\n".join([" ".join([str(x) for x in row]) for row in self.maze])

        text_foot = "\n".join([f"{P} - персонаж", f"{W} - стена", f"{F} - выход"])

        text_full = "\n\n".join([text_head, text_maze, text_foot])

        return self._maze_encoder(text_full)
