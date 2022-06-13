package db

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/mirror520/quiz/model/comment"
)

type commentRepositoryTestSuite struct {
	suite.Suite
	repo      comment.Repository
	testData1 *comment.Comment
	testData2 *comment.Comment
}

func (suite *commentRepositoryTestSuite) SetupSuite() {
	suite.repo = NewCommentRepository()

	suite.testData1 = &comment.Comment{
		UUID:     "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		ParentID: "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
		Comment:  "根據中央氣象局地震測報中心地震報告，這起規模...",
		Author:   "氣象局網站",
		Update:   time.Date(2022, 6, 1, 1, 12, 33, 0, time.Local),
		Favorite: false,
	}
	suite.repo.Store(suite.testData1)

	suite.testData2 = &comment.Comment{
		UUID:     "",
		ParentID: "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
		Comment:  "根據中央氣象局地震測報中心地震報告，這起規模...",
		Author:   "氣象局網站",
		Update:   time.Date(2022, 6, 1, 1, 12, 33, 0, time.Local),
		Favorite: false,
	}
	suite.repo.Store(suite.testData2)
}

func (suite *commentRepositoryTestSuite) TestStore() {
	c := &comment.Comment{
		UUID:     "",
		ParentID: "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
		Comment:  "根據中央氣象局地震測報中心地震報告，這起規模...",
		Author:   "氣象局網站",
		Update:   time.Date(2022, 6, 1, 1, 12, 33, 0, time.Local),
		Favorite: false,
	}

	err := suite.repo.Store(c)
	suite.NoError(err)
	suite.Len(c.UUID, 36, "UUID 長度應為 36 個字元")
}

func (suite *commentRepositoryTestSuite) TestFindCommentByUUID() {
	c, err := suite.repo.FindCommentByUUID(suite.testData1.UUID)
	suite.NoError(err)
	suite.Equal(suite.testData1.UUID, c.UUID)
}

func (suite *commentRepositoryTestSuite) TestRemove() {
	err := suite.repo.Remove(suite.testData2.UUID)
	suite.NoError(err)

	_, err = suite.repo.FindCommentByUUID(suite.testData2.UUID)
	suite.ErrorIs(err, gorm.ErrRecordNotFound)
}

func (suite *commentRepositoryTestSuite) TearDownSuite() {
	os.Remove("comment.db")
}

func TestCommentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(commentRepositoryTestSuite))
}
