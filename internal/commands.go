package internal

// "github.com/vmihailenco/msgpack/v5"

type SetObject struct {
	Type   string `json:"type" msgpack:"type"`
	Object Object `json:"object" msgpack:"object"`
	Path   string `json:"path" msgpack:"path"`
}

type SetTransform struct {
	Type   string    `json:"type" msgpack:"type"`
	Matrix []float32 `json:"matrix" msgpack:"matrix"`
	Path   string    `json:"path" msgpack:"path"`
}

type CaptureImage struct {
	Type string `json:"type" msgpack:"type"`
	Xres int    `json:"xres" msgpack:"xres"`
	Yres int    `json:"yres" msgpack:"yres"`
}

type Delete struct {
	Type string `json:"type" msgpack:"type"`
	Path string `json:"path" msgpack:"path"`
}

type SetProperty struct {
	Type        string
	Path        string
	SetProperty string
	Value       interface{}
}

type AnimationOptions struct {
	Play        bool
	Repetitions int
}

type SetAnimation struct {
	Type       string
	Animations interface{}
	Options    AnimationOptions
	Path       string
}

/*
from .geometry import Geometry, Object, Mesh, MeshPhongMaterial, OrthographicCamera, PerspectiveCamera, PointsMaterial, Points, TextTexture
from .path import Path


class SetObject:
    __slots__ = ["object", "path"]
    def __init__(self, geometry_or_object, material=None, path=None):
        if isinstance(geometry_or_object, Object):
            if material is not None:
                raise(ValueError("Please supply either an Object OR a Geometry and a Material"))
            self.object = geometry_or_object
        elif isinstance(geometry_or_object, (OrthographicCamera, PerspectiveCamera)):
            self.object = geometry_or_object
        else:
            if material is None:
                material = MeshPhongMaterial()
            if isinstance(material, PointsMaterial):
                self.object = Points(geometry_or_object, material)
            else:
                self.object = Mesh(geometry_or_object, material)
        if path is not None:
            self.path = path
        else:
            self.path = Path()

    def lower(self):
        return {
            u"type": u"set_object",
            u"object": self.object.lower(),
            u"path": self.path.lower()
        }


class SetTransform:
    __slots__ = ["matrix", "path"]
    def __init__(self, matrix, path):
        self.matrix = matrix
        self.path = path

    def lower(self):
        return {
            u"type": u"set_transform",
            u"path": self.path.lower(),
            u"matrix": list(self.matrix.T.flatten())
        }


class SetCamTarget:
    """Set the camera target point."""
    __slots__ = ["value"]
    def __init__(self, pos):
        self.value = pos

    def lower(self):
        return {
            u"type": "set_target",
            u"path": "",
            u"value": list(self.value)
        }


class CaptureImage:

    def __init__(self, xres=None, yres=None):
        self.xres = xres
        self.yres = yres

    def lower(self):
        data = {
            u"type": u"capture_image"
        }
        if self.xres:
            data[u"xres"] = self.xres
        if self.yres:
            data[u"yres"] = self.yres
        return data


class Delete:
    __slots__ = ["path"]
    def __init__(self, path):
        self.path = path

    def lower(self):
        return {
            u"type": u"delete",
            u"path": self.path.lower()
        }

class SetProperty:
    __slots__ = ["path", "key", "value"]
    def __init__(self, key, value, path):
        self.key = key
        self.value = value
        self.path = path

    def lower(self):
        return {
            u"type": u"set_property",
            u"path": self.path.lower(),
            u"property": self.key.lower(),
            u"value": self.value
        }

class SetAnimation:
    __slots__ = ["animation", "play", "repetitions"]

    def __init__(self, animation, play=True, repetitions=1):
        self.animation = animation
        self.play = play
        self.repetitions = repetitions

    def lower(self):
        return {
            u"type": u"set_animation",
            u"animations": self.animation.lower(),
            u"options": {
                u"play": self.play,
                u"repetitions": self.repetitions
            },
            u"path": ""
        }
*/

// type SceneElement struct {
// 	uid string `msgpack: "uuid"`
// }

// type Metadata struct {
// 	version   string
// 	_type     string `msgpack:"type"`
// 	generator string
// }
// type ThreeObject struct {
// 	uid      uuid.UUID `msgpack:"uuid"`
// 	name     string    `msgpacl:"name,omitempty"`
// 	_type    string    `msgpack:"type"`
// 	geometry string
// 	material string
// 	matrix   []float32
// 	children []ThreeObject `msgpack:"children,omitempty"`
// }