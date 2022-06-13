package quiz

import (
	"context"
	"testing"
	"time"

	"github.com/mirror520/quiz/model/comment"
	"github.com/mirror520/quiz/persistent/db"

	"github.com/stretchr/testify/suite"
)

type commentTestSuite struct {
	suite.Suite
	svc       Service
	repo      comment.Repository
	testData1 *comment.Comment
	testData2 *comment.Comment
	testData3 *comment.Comment
}

func (suite *commentTestSuite) SetupSuite() {
	repo := db.NewCommentRepository()
	suite.svc = NewService(repo)

	suite.repo = repo
}

func (suite *commentTestSuite) SetupTest() {
	// 測試資料應分開，每個單元測試是獨立的，不得期望測試案例執行順序
	testData1 := &comment.Comment{
		UUID:     "",
		ParentID: "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
		Comment:  "根據中央氣象局地震測報中心地震報告，這起規模...",
		Author:   "氣象局網站",
		Update:   time.Date(2022, 6, 1, 1, 12, 33, 0, time.Local),
		Favorite: false,
	}
	suite.repo.Store(testData1)

	testData2 := &comment.Comment{
		UUID:     "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		ParentID: "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
		Comment:  "根據中央氣象局地震測報中心地震報告，這起規模...",
		Author:   "氣象局網站",
		Update:   time.Date(2022, 6, 1, 1, 12, 33, 0, time.Local),
		Favorite: false,
	}
	suite.repo.Store(testData2)

	testData3 := &comment.Comment{
		UUID:     "",
		ParentID: "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
		Comment:  "根據中央氣象局地震測報中心地震報告，這起規模...",
		Author:   "氣象局網站",
		Update:   time.Date(2022, 6, 1, 1, 12, 33, 0, time.Local),
		Favorite: false,
	}
	suite.repo.Store(testData3)
}

func (suite *commentTestSuite) TestCreateComment() {
	comment := &comment.Comment{
		UUID:     "",
		ParentID: "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
		Comment:  "根據中央氣象局地震測報中心地震報告，這起規模...",
		Author:   "氣象局網站",
		Update:   time.Date(2022, 6, 1, 1, 12, 33, 0, time.Local),
		Favorite: false,
	}

	comment, err := suite.svc.CreateComment(context.TODO(), comment)
	suite.NoError(err)
	suite.Len(comment.UUID, 36, "UUID 長度應為 36 個字元")
}

func (suite *commentTestSuite) TestGetCommentByUUID() {
	comment, err := suite.svc.GetCommentByUUID(context.TODO(), suite.testData1.UUID)
	suite.NoError(err)
	suite.Equal(suite.testData1.UUID, comment.UUID, "UUID 應相同")
}

func (suite *commentTestSuite) TestModifyCommentByUUID() {
	input := &comment.Comment{
		UUID:     "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		ParentID: "a1205dab-824a-4e3a-bcd2-ed6102e60ae9",
		Comment:  "根據中央氣象局地震測報中心地震報告，這起規模...",
		Author:   "氣象局網站",
		Update:   time.Date(2022, 6, 1, 1, 12, 33, 0, time.Local),
		Favorite: true, // modify
	}

	output, err := suite.svc.ModifyCommentByUUID(context.TODO(), input, input.UUID)
	suite.NoError(err)
	suite.Equal(input.Favorite, output.Favorite, "修改後資料應與輸入資料相同")
	suite.NotEqual(suite.testData2.Favorite, output.Favorite, "修改後資料應與修改前測試資料不同")
}

func (suite *commentTestSuite) TestRemoveCommentByUUID() {
	err := suite.svc.RemoveCommentByUUID(context.TODO(), suite.testData3.UUID)
	suite.NoError(err)

	_, err = suite.svc.GetCommentByUUID(context.TODO(), suite.testData3.UUID)
	suite.ErrorIs(err, comment.ErrCommentUUIDNotFound, "確定找不到該評論")
}

func TestCommentTestSuite(t *testing.T) {
	suite.Run(t, new(commentTestSuite))
}
