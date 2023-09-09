package mysql

import (
	"ISPS/log"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func (u *UserDB) SelectConcernPeople(cp *ConcernPeople, PeopleId int) bool {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", cp.AccountId).Find(&cp).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	if (PeopleId-1)/8+1 > len(cp.ConcernPeopleID)/3 {
		return false
	}
	cpids := cp.ConcernPeopleID[((PeopleId-1)/8+1)*3-3 : ((PeopleId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if 1<<((PeopleId-1)%8)&s != 0 {
		return true
	}
	return false
}

/*查询所有关注人ID*/
func (u *UserDB) SelectAllConcernPeople(cp *ConcernPeople) []int {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", cp.AccountId).Find(&cp).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	people := make([]int, cp.ConcernPeopleNumber)
	sum := 0
	for i := 0; i < len(cp.ConcernPeopleID)/3; i++ {
		s := cp.ConcernPeopleID[3*i : 3*i+3]
		i2, err2 := strconv.Atoi(s)
		if err2 != nil {
			log1.Error(err2.Error(), "")
		}
		for j := 1; j <= 8 && i2 > 0; j++ {
			if i2&1 != 0 {
				people[sum] = 8*i + j
				sum++
			}
			i2 >>= 1
		}
	}
	return people
}

func (u *UserDB) InsertConcernPeople(cp *ConcernPeople, PeopleId int) int {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", cp.AccountId).Find(&cp).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	l := (PeopleId-1)/8 + 1
	cpid_len := len(cp.ConcernPeopleID) / 3
	where := 1 << ((PeopleId - 1) % 8)
	if l > cpid_len {
		for i := 1; i < l-cpid_len; i++ {
			cp.ConcernPeopleID = cp.ConcernPeopleID + "000"
		}
		s2 := strconv.Itoa(where)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		cp.ConcernPeopleID = cp.ConcernPeopleID + s2
		cp.ConcernPeopleNumber = cp.ConcernPeopleNumber + 1
		err := u.db.Where("account_id=?", cp.AccountId).Updates(&cp).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return 2
	} else {
		cpids := cp.ConcernPeopleID[((PeopleId-1)/8+1)*3-3 : ((PeopleId-1)/8+1)*3]
		s, err := strconv.Atoi(cpids)
		if err != nil {
			log1.Error(err.Error(), "")
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
			cp.ConcernPeopleID = cp.ConcernPeopleID[0:((PeopleId-1)/8+1)*3-3] + s2 + cp.ConcernPeopleID[((PeopleId-1)/8+1)*3:]
			cp.ConcernPeopleNumber = cp.ConcernPeopleNumber + 1
			err := u.db.Where("account_id=?", cp.AccountId).Updates(&cp).Error
			if err != nil {
				log1.Error(err.Error(), "")
			}
			return 2
		}
	}
}

func (u *UserDB) DeleteConcernPeople(cp *ConcernPeople, PeopleId int) bool {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", cp.AccountId).Find(&cp).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if (PeopleId-1)/8+1 > len(cp.ConcernPeopleID)/3 {
		return false //不存在
	}
	cpids := cp.ConcernPeopleID[((PeopleId-1)/8+1)*3-3 : ((PeopleId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if 1<<((PeopleId-1)%8)&s != 0 {
		s = s & (^(1 << ((PeopleId - 1) % 8)))
		//fmt.Println(s)
		s2 := strconv.Itoa(s)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		cp.ConcernPeopleID = cp.ConcernPeopleID[0:((PeopleId-1)/8+1)*3-3] + s2 + cp.ConcernPeopleID[((PeopleId-1)/8+1)*3:]
		cp.ConcernPeopleNumber = cp.ConcernPeopleNumber - 1
		err := u.db.Where("account_id=?", cp.AccountId).Updates(&cp).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return true
	}
	return false
}

func (u *UserDB) SelectConcernPlate(cp *ConcernPlates, PlateId int) bool {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", cp.AccountId).Find(&cp).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	if (PlateId-1)/8+1 > len(cp.ConcernPlatesID)/3 {
		return false
	}
	cpids := cp.ConcernPlatesID[((PlateId-1)/8+1)*3-3 : ((PlateId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if 1<<((PlateId-1)%8)&s != 0 {
		return true
	}
	return false
}

/*查询所有关注板块ID*/
func (u *UserDB) SelectAllConcernPlates(cp *ConcernPlates) []int {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", cp.AccountId).Find(&cp).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	plates := make([]int, cp.ConcernPlatesNumber)
	sum := 0
	for i := 0; i < len(cp.ConcernPlatesID)/3; i++ {
		s := cp.ConcernPlatesID[3*i : 3*i+3]
		i2, err2 := strconv.Atoi(s)
		if err2 != nil {
			log1.Error(err2.Error(), "")
		}
		for j := 1; j <= 8 && i2 > 0; j++ {
			if i2&1 != 0 {
				plates[sum] = 8*i + j
				sum++
			}
			i2 >>= 1
		}
	}
	return plates
}

func (u *UserDB) InsertConcernPlate(cp *ConcernPlates, PlateId int) int {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", cp.AccountId).Find(&cp).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	l := (PlateId-1)/8 + 1
	cpid_len := len(cp.ConcernPlatesID) / 3
	where := 1 << ((PlateId - 1) % 8)
	if l > cpid_len {
		for i := 1; i < l-cpid_len; i++ {
			cp.ConcernPlatesID = cp.ConcernPlatesID + "000"
		}
		s2 := strconv.Itoa(where)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		cp.ConcernPlatesID = cp.ConcernPlatesID + s2
		cp.ConcernPlatesNumber = cp.ConcernPlatesNumber + 1
		err := u.db.Where("account_id=?", cp.AccountId).Updates(&cp).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return 2
	} else {
		cpids := cp.ConcernPlatesID[((PlateId-1)/8+1)*3-3 : ((PlateId-1)/8+1)*3]
		s, err := strconv.Atoi(cpids)
		fmt.Println(s)
		if err != nil {
			log1.Error(err.Error(), "")
		}
		if where&s != 0 {
			return 1 //已存在
		} else {
			s = s | where
			fmt.Println(s)
			s2 := strconv.Itoa(s)
			if len(s2) == 1 {
				s2 = "00" + s2
			} else if len(s2) == 2 {
				s2 = "0" + s2
			}
			cp.ConcernPlatesID = cp.ConcernPlatesID[0:((PlateId-1)/8+1)*3-3] + s2 + cp.ConcernPlatesID[((PlateId-1)/8+1)*3:]
			cp.ConcernPlatesNumber = cp.ConcernPlatesNumber + 1
			err1 := u.db.Where("account_id=?", cp.AccountId).Updates(&cp).Error
			if err1 != nil {
				log1.Error(err1.Error(), "")
			}
			return 2
		}
	}
}

func (u *UserDB) DeleteConcernPlates(cp *ConcernPlates, PlateId int) bool {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", cp.AccountId).Find(&cp).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if (PlateId-1)/8+1 > len(cp.ConcernPlatesID)/3 {
		return false //不存在
	}
	cpids := cp.ConcernPlatesID[((PlateId-1)/8+1)*3-3 : ((PlateId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if 1<<((PlateId-1)%8)&s != 0 {
		s = s & (^(1 << ((PlateId - 1) % 8)))
		//fmt.Println(s)
		s2 := strconv.Itoa(s)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		cp.ConcernPlatesID = cp.ConcernPlatesID[0:((PlateId-1)/8+1)*3-3] + s2 + cp.ConcernPlatesID[((PlateId-1)/8+1)*3:]
		cp.ConcernPlatesNumber = cp.ConcernPlatesNumber - 1
		err := u.db.Where("account_id=?", cp.AccountId).Updates(&cp).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return true
	}
	return false
}

func (u *UserDB) SelectUserFans(uf *UserFans, FansId int) bool {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", uf.AccountId).Find(&uf).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	if (FansId-1)/8+1 > len(uf.FansID)/3 {
		return false
	}
	ufids := uf.FansID[((FansId-1)/8+1)*3-3 : ((FansId-1)/8+1)*3]
	s, err := strconv.Atoi(ufids)
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if 1<<((FansId-1)%8)&s != 0 {
		return true
	}
	return false
}

/*查询所有用户粉丝ID*/
func (u *UserDB) SelectAllUserFans(uf *UserFans) []int {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", uf.AccountId).Find(&uf).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	userfans := make([]int, uf.FansNumber)
	sum := 0
	for i := 0; i < len(uf.FansID)/3; i++ {
		s := uf.FansID[3*i : 3*i+3]
		i2, err2 := strconv.Atoi(s)
		if err2 != nil {
			log1.Error(err2.Error(), "")
		}
		for j := 1; j <= 8 && i2 > 0; j++ {
			if i2&1 != 0 {
				userfans[sum] = 8*i + j
				sum++
			}
			i2 >>= 1
		}
	}
	return userfans
}

func (u *UserDB) InsertFans(uf *UserFans, FansId int) int {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", uf.AccountId).Find(&uf).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	l := (FansId-1)/8 + 1
	cpid_len := len(uf.FansID) / 3
	where := 1 << ((FansId - 1) % 8)
	if l > cpid_len {
		for i := 1; i < l-cpid_len; i++ {
			uf.FansID = uf.FansID + "000"
		}
		s2 := strconv.Itoa(where)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		uf.FansID = uf.FansID + s2
		uf.FansNumber = uf.FansNumber + 1
		err := u.db.Where("account_id=?", uf.AccountId).Updates(&uf).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return 2
	} else {
		ufids := uf.FansID[((FansId-1)/8+1)*3-3 : ((FansId-1)/8+1)*3]
		s, err := strconv.Atoi(ufids)
		if err != nil {
			log1.Error(err.Error(), "")
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
			uf.FansID = uf.FansID[0:((FansId-1)/8+1)*3-3] + s2 + uf.FansID[((FansId-1)/8+1)*3:]
			uf.FansNumber = uf.FansNumber + 1
			err1 := u.db.Where("account_id=?", uf.AccountId).Updates(&uf).Error
			if err1 != nil {
				log1.Error(err1.Error(), "")
			}
			return 2
		}
	}
}

func (u *UserDB) DeleteFans(uf *UserFans, FansId int) bool {
	log1 := log.NewLog()
	u.db = Init("user")
	err := u.db.Where("account_id", uf.AccountId).Find(&uf).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if (FansId-1)/8+1 > len(uf.FansID)/3 {
		return false //不存在
	}
	cpids := uf.FansID[((FansId-1)/8+1)*3-3 : ((FansId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if 1<<((FansId-1)%8)&s != 0 {
		s = s & (^(1 << ((FansId - 1) % 8)))
		//fmt.Println(s)
		s2 := strconv.Itoa(s)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		uf.FansID = uf.FansID[0:((FansId-1)/8+1)*3-3] + s2 + uf.FansID[((FansId-1)/8+1)*3:]
		uf.FansNumber = uf.FansNumber - 1
		err := u.db.Where("account_id=?", uf.AccountId).Updates(&uf).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return true
	}
	return false
}
