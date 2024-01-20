package converter

import (
	"auth/internal/model"
	dbModel "auth/internal/repository/user/model"
	"database/sql"
	"time"
)

func ToUserFromRepo(repoUser dbModel.User) *model.User {
	return &model.User{
		UUID:      repoUser.UUID,
		Profile:   *ToUserProfileFromRepo(repoUser.Profile),
		CreatedAt: repoUser.CreatedAt,
		UpdatedAt: sqlNullTimeToTime(repoUser.UpdatedAt),
	}
}

func ToUserProfileFromRepo(repoUserProfile dbModel.UserProfile) *model.UserProfile {
	return &model.UserProfile{
		FirstName: repoUserProfile.FirstName,
		LastName:  repoUserProfile.LastName,
		Age:       repoUserProfile.Age,
		CreatedAt: repoUserProfile.CreatedAt,
		UpdatedAt: sqlNullTimeToTime(repoUserProfile.UpdatedAt),
	}
}

func ToRepoUserFromUser(user *model.User) *dbModel.User {
	return &dbModel.User{
		UUID:      user.UUID,
		Profile:   *ToRepoUserProfileFromUser(&user.Profile),
		CreatedAt: user.CreatedAt,
		UpdatedAt: timeToSqlNullTime(user.UpdatedAt),
	}
}

func ToRepoUserProfileFromUser(userProfile *model.UserProfile) *dbModel.UserProfile {
	return &dbModel.UserProfile{
		FirstName: userProfile.FirstName,
		LastName:  userProfile.LastName,
		Age:       userProfile.Age,
		CreatedAt: userProfile.CreatedAt,
		UpdatedAt: timeToSqlNullTime(userProfile.UpdatedAt),
	}
}

func sqlNullTimeToTime(sqlNullTime sql.NullTime) *time.Time {
	if sqlNullTime.Valid {
		return &sqlNullTime.Time
	}
	return nil
}

func timeToSqlNullTime(time *time.Time) sql.NullTime {
	if time == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Valid: true,
		Time:  *time,
	}
}
