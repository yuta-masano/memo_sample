package db

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func getTagRepositoryForTest() *TagRepository {
	return NewTagRepository(db)
}

func TestTagSaveInDBSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getTagRepositoryForTest()

	ctx := context.Background()

	// 1件名
	_, err := repo.Save(ctx, "Tag First")
	if err != nil {
		t.Error("failed TestTagSaveInTagrySuccess Save", err)
	}

}

func TestTagTransactionCommitSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repo := getTagRepositoryForTest()

	ctx := context.Background()

	ctx, err := repo.Begin(ctx)
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	_, err = repo.Save(ctx, "Transaction Commit Test")
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}

	_, err = repo.Commit(ctx)
	if err != nil {
		repo.Rollback(ctx)
		panic(err)
	}
}

func TestTagAndMemoTransactionCommitSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()

	ctx := context.Background()

	ctx, err := repoM.Begin(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	memo, err := repoM.Save(ctx, "Transaction Commit Memo")
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	tag, err := repoT.Save(ctx, "Transaction Commit Tag")
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}

	err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}

	_, err = repoM.Commit(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}
}

func TestTagAndMemoTransactionRollbackSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()

	ctx := context.Background()

	ctx, err := repoM.Begin(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	memo, err := repoM.Save(ctx, "Transaction Rollback Memo")
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	tag, err := repoT.Save(ctx, "Transaction Rollback Tag")
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}

	err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}

	// 強制的にロールバック
	repoM.Rollback(ctx)
}

func TestTagAndMemoGetAllByMemoIDSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()

	ctx := context.Background()

	ctx, err := repoM.Begin(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	memo, err := repoM.Save(ctx, "GetAllByMemoID Test Memo")
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	tag, err := repoT.Save(ctx, "GetAllByMemoID Test Tag")
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}

	err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}

	ctx, err = repoM.Commit(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	flag := false
	list, err := repoT.GetAllByMemoID(ctx, memo.ID)
	for _, tg := range list {
		if tg.ID == tag.ID {
			flag = true
		}
	}

	if !flag {
		panic(fmt.Errorf("GetAllByMemoID Error"))
	}
}

func TestTagAndMemoSearchMemoIDsByTitleSuccess(t *testing.T) {
	connectTestDB()
	defer closeTestDB()

	repoT := getTagRepositoryForTest()
	repoM := getMemoRepositoryForTest()

	ctx := context.Background()

	ctx, err := repoM.Begin(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	memo, err := repoM.Save(ctx, "SearchMemoIDsByTitle Test Memo")
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	tag, err := repoT.Save(ctx, "SearchMemoIDsByTitle Test Tag")
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}

	err = repoT.SaveTagAndMemo(ctx, tag.ID, memo.ID)
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}

	tag2, err := repoT.Get(ctx, tag.ID)
	if err != nil {
		repoT.Rollback(ctx)
		panic(err)
	}
	t.Log(tag2)

	ctx, err = repoM.Commit(ctx)
	if err != nil {
		repoM.Rollback(ctx)
		panic(err)
	}

	flag := false
	list, err := repoT.SearchMemoIDsByTitle(ctx, tag.Title)
	for _, id := range list {
		if id == memo.ID {
			flag = true
		}
	}

	if !flag {
		panic(fmt.Errorf("SearchMemoIDsByTitle Error"))
	}
}