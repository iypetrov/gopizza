package dtos

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/internal/toasts"
)

type SaladRequest struct {
	Name     string
	Tomatoes bool
	Garlic   bool
	Onion    bool
	Parmesan bool
	Chicken  bool
	Image    io.Reader
	Price    float64
}

func (req *SaladRequest) Validate() map[string]string {
	errs := make(map[string]string)

	if len(req.Name) == 0 {
		errs["name"] = toasts.ErrNameIsRequired.Error()
	}

	if req.Image == nil {
		errs["image"] = toasts.ErrImageIsRequired.Error()
	} else {
		buf := make([]byte, 1)
		if _, err := req.Image.Read(buf); err == io.EOF {
			errs["image"] = toasts.ErrImageIsRequired.Error()
		} else {
			if seeker, ok := req.Image.(io.Seeker); ok {
				seeker.Seek(0, io.SeekStart)
			}
		}
	}

	if req.Price == 0 {
		errs["price"] = toasts.ErrPriceShouldBePositiveNumber.Error()
	}

	return errs
}

func ParseToSaladRequest(r *http.Request) (SaladRequest, error) {
	// 10 MB limit
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return SaladRequest{}, err
	}

	var req SaladRequest
	req.Name = parseString(r, "name")
	req.Tomatoes = parseBool(r, "tomatoes")
	req.Garlic = parseBool(r, "garlic")
	req.Onion = parseBool(r, "onion")
	req.Parmesan = parseBool(r, "parmesan")
	req.Chicken = parseBool(r, "chicken")
	req.Price = parseFloat(r, "price")
	file, _, err := r.FormFile("image")
	if err != nil {
		return SaladRequest{}, err
	}
	defer file.Close()
	req.Image = file

	return req, nil
}

type SaladResponse struct {
	ID        uuid.UUID
	Name      string
	Tomatoes  bool
	Garlic    bool
	Onion     bool
	Parmesan  bool
	Chicken   bool
	ImageUrl  string
	Price     float64
	UpdatedAt time.Time
}

func (req *SaladResponse) Description() string {
	var ingredients []string

	if req.Tomatoes {
		ingredients = append(ingredients, "tomatoes")
	}
	if req.Garlic {
		ingredients = append(ingredients, "garlic")
	}
	if req.Onion {
		ingredients = append(ingredients, "onion")
	}
	if req.Parmesan {
		ingredients = append(ingredients, "parmesan")
	}
	if req.Chicken {
		ingredients = append(ingredients, "chicken")
	}

	description := strings.Join(ingredients, ", ")
	if len(description) == 0 {
		return description
	}
	return strings.ToUpper(string(description[0])) + description[1:]
}
