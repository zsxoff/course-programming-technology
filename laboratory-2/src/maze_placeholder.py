import numpy as np

E = 0
W = 1
F = 2
P = 3

MAZE_MAP_CLASSIC = np.asarray(
    [
        [W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
        [W, P, E, E, E, E, E, E, E, E, E, E, E, E, E, E, E, F],
        [W, E, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W, W],
        [W, E, W, W, E, W, W, W, W, W, W, W, W, W, W, W, W, W],
        [W, E, E, E, E, W, W, W, W, W, W, W, W, W, W, W, W, W],
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
