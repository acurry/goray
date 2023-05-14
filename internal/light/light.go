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
	position *point.Point
	ambient  *point.Point
	diffuse  *point.Point
	specular *point.Point
}

func New() *Light {
	return &Light{
		point.New(),
		point.New(),
		point.New(),
		point.New(),
	}
}
