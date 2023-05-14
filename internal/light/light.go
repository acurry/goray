package light

import "github.com/acurry/goray/internal/point"

/*
@dataclass
class Light:
    position: list[float] = field(default_factory=list)
    ambient: list[float] = field(default_factory=list)
    diffuse: list[float] = field(default_factory=list)
    specular: list[float] = field(default_factory=list)

    def __post_init__(self):
        self.position = np.array(self.position)
        self.ambient = np.array(self.ambient)
        self.diffuse = np.array(self.diffuse)
        self.specular = np.array(self.specular)
*/

type Light struct {
	Position point.Point `yaml:",flow"`
	Ambient  point.Point `yaml:",flow"`
	Diffuse  point.Point `yaml:",flow"`
	Specular point.Point `yaml:",flow"`
}
