package mysql

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type PermissionDB struct {
	db *gorm.DB
}

func (p *PermissionDB) SelectManage(m *Manage, AccountId int) bool {
	p.db = Init("permission")
	err := p.db.Where("plate_id", m.PlateId).Find(&m).Error
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Println(cp.ConcernPlatesID)
	if (AccountId-1)/8+1 > len(m.AccountId)/3 {
		return false
	}
	cpids := m.AccountId[((AccountId-1)/8+1)*3-3 : ((AccountId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		fmt.Println(err)
	}
	if 1<<((AccountId-1)%8)&s != 0 {
		return true
	}
	return false
}

func (p *PermissionDB) InsertManage(m *Manage, AccountId int) int {
	p.db = Init("permission")
	err := p.db.Where("plate_id", m.PlateId).Find(&m).Error
	if err != nil {
		fmt.Print(err)
	}
	l := (AccountId-1)/8 + 1
	cpid_len := len(m.AccountId) / 3
	where := 1 << ((AccountId - 1) % 8)
	if l > cpid_len {
		for i := 1; i < l-cpid_len; i++ {
			m.AccountId = m.AccountId + "000"
		}
		s2 := strconv.Itoa(where)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		m.AccountId = m.AccountId + s2
		err := p.db.Where("plate_id=?", m.PlateId).Updates(&m).Error
		if err != nil {
			fmt.Print(err)
		}
		return 2
	} else {
		cpids := m.AccountId[((AccountId-1)/8+1)*3-3 : ((AccountId-1)/8+1)*3]
		s, err := strconv.Atoi(cpids)
		if err != nil {
			fmt.Println(err)
		}
		if where&s != 0 {
			return 1 //已存在
		} else {
			s = s | where
			s2 := strconv.Itoa(s)
			if len(s2) == 1 {
				s2 = "00" + s2
			} else if len(s2) == 2 {
				s2 = "0" + s2
			}
			m.AccountId = m.AccountId[0:((AccountId-1)/8+1)*3-3] + s2 + m.AccountId[((AccountId-1)/8+1)*3:]
			err1 := p.db.Where("plate_id=?", m.PlateId).Updates(&m).Error
			if err1 != nil {
				fmt.Print(err1)
			}
			return 2
		}
	}
}

func (p *PermissionDB) DeleteManage(m *Manage, AccountId int) bool {
	p.db = Init("permission")
	err := p.db.Where("plate_id", m.PlateId).Find(&m).Error
	if err != nil {
		fmt.Print(err)
	}
	if (AccountId-1)/8+1 > len(m.AccountId)/3 {
		return false //不存在
	}
	cpids := m.AccountId[((AccountId-1)/8+1)*3-3 : ((AccountId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		fmt.Println(err)
	}
	if 1<<((AccountId-1)%8)&s != 0 {
		s = s & (^(1 << ((AccountId - 1) % 8)))
		//fmt.Println(s)
		s2 := strconv.Itoa(s)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		m.AccountId = m.AccountId[0:((AccountId-1)/8+1)*3-3] + s2 + m.AccountId[((AccountId-1)/8+1)*3:]
		err := p.db.Where("plate_id=?", m.PlateId).Updates(&m).Error
		if err != nil {
			fmt.Print(err)
		}
		return true
	}
	return false
}

func (p *PermissionDB) SelectModerators(m *Moderators, PlatesId int) bool {
	p.db = Init("permission")
	err := p.db.Where("account_id", m.AccountId).Find(&m).Error
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Println(cp.ConcernPlatesID)
	if (PlatesId-1)/8+1 > len(m.PlatesId)/3 {
		return false
	}
	cpids := m.PlatesId[((PlatesId-1)/8+1)*3-3 : ((PlatesId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		fmt.Println(err)
	}
	if 1<<((PlatesId-1)%8)&s != 0 {
		return true
	}
	return false
}

func (p *PermissionDB) InsertModerators(m *Moderators, PlatesId int) int {
	p.db = Init("permission")
	err := p.db.Where("account_id", m.AccountId).Find(&m).Error
	if err != nil {
		fmt.Print(err)
	}
	l := (PlatesId-1)/8 + 1
	cpid_len := len(m.PlatesId) / 3
	where := 1 << ((PlatesId - 1) % 8)
	if l > cpid_len {
		for i := 1; i < l-cpid_len; i++ {
			m.PlatesId = m.PlatesId + "000"
		}
		s2 := strconv.Itoa(where)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		m.PlatesId = m.PlatesId + s2
		err := p.db.Where("account_id", m.AccountId).Updates(&m).Error
		if err != nil {
			fmt.Print(err)
		}
		return 2
	} else {
		cpids := m.PlatesId[((PlatesId-1)/8+1)*3-3 : ((PlatesId-1)/8+1)*3]
		s, err := strconv.Atoi(cpids)
		if err != nil {
			fmt.Println(err)
		}
		if where&s != 0 {
			return 1 //已存在
		} else {
			s = s | where
			s2 := strconv.Itoa(s)
			if len(s2) == 1 {
				s2 = "00" + s2
			} else if len(s2) == 2 {
				s2 = "0" + s2
			}
			m.PlatesId = m.PlatesId[0:((PlatesId-1)/8+1)*3-3] + s2 + m.PlatesId[((PlatesId-1)/8+1)*3:]
			err1 := p.db.Where("account_id=?", m.AccountId).Updates(&m).Error
			if err1 != nil {
				fmt.Print(err1)
			}
			return 2
		}
	}
}

func (p *PermissionDB) DeleteModerators(m *Moderators, PlatesId int) bool {
	p.db = Init("permission")
	err := p.db.Where("account_id", m.AccountId).Find(&m).Error
	if err != nil {
		fmt.Print(err)
	}
	if (PlatesId-1)/8+1 > len(m.PlatesId)/3 {
		return false //不存在
	}
	cpids := m.PlatesId[((PlatesId-1)/8+1)*3-3 : ((PlatesId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		fmt.Println(err)
	}
	if 1<<((PlatesId-1)%8)&s != 0 {
		s = s & (^(1 << ((PlatesId - 1) % 8)))
		//fmt.Println(s)
		s2 := strconv.Itoa(s)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		m.PlatesId = m.PlatesId[0:((PlatesId-1)/8+1)*3-3] + s2 + m.PlatesId[((PlatesId-1)/8+1)*3:]
		err := p.db.Where("account_id=?", m.AccountId).Updates(&m).Error
		if err != nil {
			fmt.Print(err)
		}
		return true
	}
	return false
}
