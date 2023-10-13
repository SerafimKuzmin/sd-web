package middleware

import (
	friendUsecase "github.com/SerafimKuzmin/sd/src/internal/List/usecase"
	echo "github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"github.com/SerafimKuzmin/sd/src/models"
)

type AclMiddleware struct {
	friendUC friendUsecase.UsecaseI
}

func NewAclMiddleware(friendUC friendUsecase.UsecaseI) *AclMiddleware {
	return &AclMiddleware{friendUC: friendUC}
}

func (am *AclMiddleware) AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*models.User)
		if user.Role != 2 {
			return echo.NewHTTPError(http.StatusForbidden, models.ErrPermissionDenied.Error())
		}
		return next(c)
	}
}

// for all handlers with c.Param("user_id")
func (am *AclMiddleware) FriendsOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authUser, ok := c.Get("user").(*models.User)

		if !ok {
			c.Logger().Error("can't get user from context")
			return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
		}

		otherUserID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
		}

		isFriends, err := am.friendUC.CheckIsFriends(authUser.ID, otherUserID)

		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
		}

		if !isFriends {
			c.Logger().Error("Error: FriendsOnly middleware it is not a friends")
			return echo.NewHTTPError(http.StatusForbidden, models.ErrPermissionDenied.Error())
		}

		return next(c)
	}
}

// for all handlers with c.Param("user_id")
func (am *AclMiddleware) FriendsOrAdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authUser, ok := c.Get("user").(*models.User)

		if !ok {
			c.Logger().Error("can't get user from context")
			return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServerError.Error())
		}

		if authUser.Role == 2 {
			return next(c)
		}

		otherUserID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
		}

		isFriends, err := am.friendUC.CheckIsFriends(authUser.ID, otherUserID)

		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadRequest.Error())
		}

		if !isFriends {
			c.Logger().Error("Error: FriendsOrAdminOnly middleware it is not a friends and you are not an admin")
			return echo.NewHTTPError(http.StatusForbidden, models.ErrPermissionDenied.Error())
		}

		return next(c)
	}
}
