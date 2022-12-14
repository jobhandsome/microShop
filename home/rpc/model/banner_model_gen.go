// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	bannerFieldNames          = builder.RawFieldNames(&Banner{})
	bannerRows                = strings.Join(bannerFieldNames, ",")
	bannerRowsExpectAutoSet   = strings.Join(stringx.Remove(bannerFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	bannerRowsWithPlaceHolder = strings.Join(stringx.Remove(bannerFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	bannerModel interface {
		Insert(ctx context.Context, data *Banner) (sql.Result, error)
		FindAll(ctx context.Context, orderBy string) ([]Banner, error)
		FindOne(ctx context.Context, id uint64) (*Banner, error)
		FindOneById(ctx context.Context, id uint64) (*Banner, error)
		Update(ctx context.Context, data *Banner) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultBannerModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Banner struct {
		Id             uint64 `db:"id"`
		Title          string `db:"title"`
		Href           string `db:"href"`
		Target         string `db:"target"`
		Pic            string `db:"pic"`
		Status         int64  `db:"status"`
		CreateTimeTamp int64  `db:"createTimeTamp"`
		UpdateTimeTamp int64  `db:"updateTimeTamp"`
	}
)

func newBannerModel(conn sqlx.SqlConn) *defaultBannerModel {
	return &defaultBannerModel{
		conn:  conn,
		table: "`banner`",
	}
}

func (m *defaultBannerModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultBannerModel) FindAll(ctx context.Context, orderBy string) ([]Banner, error) {
	query := fmt.Sprintf("select %s from %s order by createTimeTamp %s", bannerRows, m.table, orderBy)
	var resp []Banner
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBannerModel) FindOne(ctx context.Context, id uint64) (*Banner, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bannerRows, m.table)
	var resp Banner
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBannerModel) FindOneById(ctx context.Context, id uint64) (*Banner, error) {
	var resp Banner
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bannerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBannerModel) Insert(ctx context.Context, data *Banner) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, bannerRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Title, data.Href, data.Target, data.Pic, data.Status, data.CreateTimeTamp, data.UpdateTimeTamp)
	return ret, err
}

func (m *defaultBannerModel) Update(ctx context.Context, newData *Banner) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bannerRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Title, newData.Href, newData.Target, newData.Pic, newData.Status, newData.CreateTimeTamp, newData.UpdateTimeTamp, newData.Id)
	return err
}

func (m *defaultBannerModel) tableName() string {
	return m.table
}
