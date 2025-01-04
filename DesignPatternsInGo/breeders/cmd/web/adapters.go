package main

// import (
// 	"breeders/models"
// 	"encoding/json"
// 	"encoding/xml"
// 	"io"
// 	"net/http"
// )

// // CatBreedsInterface is simply our target interface, which defines all the
// // methods that any type which implements this interface must have.
// type CatBreedsInterface interface {
// 	GetAllCatBreeds() ([]*models.CatBreed, error)
// }

// // RemoteService is the Adapter type. It embeds a DataInterface interface
// // (which is critical to the pattern).
// type RemoteService struct {
// 	Remote CatBreedsInterface
// }

// // GetAllBreeds is the function on RemoteService which lets us call any
// // adapter which implements the DataInterface type.
// func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
// 	return rs.Remote.GetAllCatBreeds()
// }

// // JSONBackend is the JSON adaptee, which needs to satisfy the CatBreedsInterface
// // by having a GetAllCatsBreeds method.
// type JSONBackend struct{}

// // GetAllCatBreeds is necessary so that JSONBackend satisfies the
// // CatBreedsInterface requirements.
// func (jb *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
// 	resp, err := http.Get("http://127.0.0.1:8081/api/cat-breeds/all/json")
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var breeds []*models.CatBreed
// 	err = json.Unmarshal(body, &breeds)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return breeds, nil
// }

// // XMLBackend is the XML adaptee, which needs to satisfy the CatBreedsInterface
// // by having a GetAllCatsBreeds method.
// type XMLBackend struct{}

// // GetAllCatBreeds is necessary so that XMLBackend satisfies the
// // CatBreedsInterface requirements.
// func (xb *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
// 	resp, err := http.Get("http://127.0.0.1:8081/api/cat-breeds/all/xml")
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	type catBreeds struct {
// 		XMLName struct{}           `xml:"cat-breeds"`
// 		Breeds  []*models.CatBreed `xml:"cat-breed"`
// 	}

// 	var breeds catBreeds

// 	err = xml.Unmarshal(body, &breeds)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return breeds.Breeds, nil
// }
