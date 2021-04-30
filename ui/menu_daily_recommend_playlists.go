package ui

import (
	"github.com/anhoder/netease-music/service"
	"go-musicfox/ds"
	"go-musicfox/utils"
)

type DailyRecommendPlaylistsMenu struct {
	menus []MenuItem
}

func (m *DailyRecommendPlaylistsMenu) BeforeBackMenuHook() Hook {
	return nil
}

func (m *DailyRecommendPlaylistsMenu) IsPlayable() bool {
	return false
}

func (m *DailyRecommendPlaylistsMenu) ResetPlaylistWhenPlay() bool {
	return false
}

func (m *DailyRecommendPlaylistsMenu) GetMenuKey() string {
	return "daily_playlists"
}

func (m *DailyRecommendPlaylistsMenu) MenuViews() []MenuItem {
	return m.menus
}

func (m *DailyRecommendPlaylistsMenu) SubMenu(model *NeteaseModel, index int) IMenu {
	playlists, ok := model.menuData.([]ds.Playlist)
	if !ok || len(playlists) < index {
		return nil
	}
	return &PlaylistDetailMenu{PlaylistId: playlists[index].Id}
}

func (m *DailyRecommendPlaylistsMenu) ExtraView() string {
	return ""
}

func (m *DailyRecommendPlaylistsMenu) BeforePrePageHook() Hook {
	// Nothing to do
	return nil
}

func (m *DailyRecommendPlaylistsMenu) BeforeNextPageHook() Hook {
	// Nothing to do
	return nil
}

func (m *DailyRecommendPlaylistsMenu) BeforeEnterMenuHook() Hook {
	return func(model *NeteaseModel) bool {
		if utils.CheckUserInfo(model.user) == utils.NeedLogin {
			NeedLoginHandle(model, enterMenu)
			return false
		}

		recommendPlaylists := service.RecommendResourceService{}
		code, response := recommendPlaylists.RecommendResource()
		codeType := utils.CheckCode(code)
		if codeType == utils.NeedLogin {
			NeedLoginHandle(model, enterMenu)
			return false
		}
		list := utils.GetDailyPlaylists(response)
		for _, playlist := range list {
			m.menus = append(m.menus, MenuItem{utils.ReplaceSpecialStr(playlist.Name), ""})
		}

		model.menuData = list

		return true
	}
}

func (m *DailyRecommendPlaylistsMenu) BottomOutHook() Hook {
	// Nothing to do
	return nil
}

func (m *DailyRecommendPlaylistsMenu) TopOutHook() Hook {
	// Nothing to do
	return nil
}

