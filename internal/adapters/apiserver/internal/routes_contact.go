package internal

import (
	"net/http"

	"github.com/Sreeram-ganesan/my-blog/internal/core/usecase"
	"github.com/gin-gonic/gin"
)

func CreateContact(uc *usecase.UseCases) func(c *gin.Context) {
	return func(c *gin.Context) {
		req := new(ContactToSaveRest)
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, NewBadRequestErrResponse(err))
			return
		}
		contactToSave, err := req.toModel()
		if err != nil {
			c.JSON(http.StatusBadRequest, NewBadRequestErrResponse(err))
			return
		}
		contact, err := uc.AddAddrBookContact(c.Request.Context(), contactToSave)
		if err != nil {
			c.JSON(http.StatusInternalServerError, NewInternalServerErrResponse(err))
			return
		}
		contactRest := contactModelToRest(contact)
		c.JSON(http.StatusCreated, contactRest)
	}
}

func UpdateContact(uc *usecase.UseCases) func(c *gin.Context) {
	return func(c *gin.Context) {
		ID := c.Param("id")
		req := new(ContactToSaveRest)
		if err := c.Bind(req); err != nil {
			c.JSON(http.StatusBadRequest, NewBadRequestErrResponse(err))
			return
		}
		contactToSave, err := req.toModel()
		if err != nil {
			c.JSON(http.StatusBadRequest, NewBadRequestErrResponse(err))
			return
		}
		contact, found, err := uc.UpdateAddrBookContact(c.Request.Context(), ID, contactToSave)
		if err != nil {
			c.JSON(http.StatusInternalServerError, NewInternalServerErrResponse(err))
			return
		}
		if !found {
			c.JSON(http.StatusNotFound, NotFoundErrResponse)
			return
		}
		contactRest := contactModelToRest(contact)
		c.JSON(http.StatusOK, contactRest)
	}
}

func ListAllContacts(uc *usecase.UseCases) func(c *gin.Context) {
	return func(c *gin.Context) {
		contacts, err := uc.LoadAddrBookContacts(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, NewInternalServerErrResponse(err))
			return
		}
		contactRestList := make([]*ContactRest, len(contacts))
		for i, contact := range contacts {
			contactRestList[i] = contactModelToRest(contact)
		}
		c.JSON(http.StatusOK, contactRestList)
	}
}

func GetContact(uc *usecase.UseCases) func(c *gin.Context) {
	return func(c *gin.Context) {
		ID := c.Param("id")
		contact, err := uc.LoadAddrBookContactByID(c.Request.Context(), ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, NewInternalServerErrResponse(err))
			return
		}
		if contact == nil {
			c.JSON(http.StatusNotFound, NotFoundErrResponse)
			return
		}
		contactRest := contactModelToRest(contact)
		c.JSON(http.StatusOK, contactRest)
	}
}

func DeleteContact(uc *usecase.UseCases) func(c *gin.Context) {
	return func(c *gin.Context) {
		ID := c.Param("id")
		found, err := uc.DeleteAddrBookContact(c.Request.Context(), ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, NewInternalServerErrResponse(err))
			return
		}
		if !found {
			c.JSON(http.StatusNotFound, NotFoundErrResponse)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

func CreateBlog(uc *usecase.UseCases) func(c *gin.Context) {
	return func(c *gin.Context) {
		req := new(BlogToSaveRest)
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, NewBadRequestErrResponse(err))
			return
		}
		blogToSave, err := req.toModel()
		if err != nil {
			c.JSON(http.StatusBadRequest, NewBadRequestErrResponse(err))
			return
		}
		blog, err := uc.AddBlog(c.Request.Context(), blogToSave)
		if err != nil {
			c.JSON(http.StatusInternalServerError, NewInternalServerErrResponse(err))
			return
		}
		blogRest := blogModelToRest(blog)
		c.JSON(http.StatusCreated, blogRest)
	}
}
