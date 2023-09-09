package mysql

import (
	"ISPS/log"
	"strconv"

	"gorm.io/gorm"
)

type PlateDB struct {
	db *gorm.DB
}

func (p *PlateDB) SelectPlateFans(pf *PlateFans, FansId int) bool {
	log1 := log.NewLog()
	p.db = Init("plate")
	err := p.db.Where("plate_id", pf.PlateId).Find(&pf).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	if (FansId-1)/8+1 > len(pf.FansID)/3 {
		return false
	}
	cpids := pf.FansID[((FansId-1)/8+1)*3-3 : ((FansId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if 1<<((FansId-1)%8)&s != 0 {
		return true
	}
	return false
}

func (p *PlateDB) SelectAllPlateFans(pf *PlateFans) []int {
	log1 := log.NewLog()
	p.db = Init("plate")
	err := p.db.Where("plate_id", pf.PlateId).Find(&pf).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	platefans := make([]int, pf.FansNumber)
	sum := 0
	for i := 0; i < len(pf.FansID)/3; i++ {
		s := pf.FansID[3*i : 3*i+3]
		i2, err2 := strconv.Atoi(s)
		if err2 != nil {
			log1.Error(err2.Error(), "")
		}
		for j := 1; j <= 8 && i2 > 0; j++ {
			if i2&1 != 0 {
				platefans[sum] = 8*i + j
				sum++
			}
			i2 >>= 1
		}
	}
	return platefans
}

func (p *PlateDB) InsertPlateFans(pf *PlateFans, FansId int) int {
	log1 := log.NewLog()
	p.db = Init("plate")
	err := p.db.Where("plate_id", pf.PlateId).Find(&pf).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	l := (FansId-1)/8 + 1
	cpid_len := len(pf.FansID) / 3
	where := 1 << ((FansId - 1) % 8)
	if l > cpid_len {
		for i := 1; i < l-cpid_len; i++ {
			pf.FansID = pf.FansID + "000"
		}
		s2 := strconv.Itoa(where)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		pf.FansID = pf.FansID + s2
		pf.FansNumber = pf.FansNumber + 1
		err := p.db.Where("plate_id=?", pf.PlateId).Updates(&pf).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return 2
	} else {
		cpids := pf.FansID[((FansId-1)/8+1)*3-3 : ((FansId-1)/8+1)*3]
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
			pf.FansID = pf.FansID[0:((FansId-1)/8+1)*3-3] + s2 + pf.FansID[((FansId-1)/8+1)*3:]
			pf.FansNumber = pf.FansNumber + 1
			err1 := p.db.Where("plate_id=?", pf.PlateId).Updates(&pf).Error
			if err1 != nil {
				log1.Error(err1.Error(), "")
			}
			return 2
		}
	}
}

func (p *PlateDB) DeletePlateFans(pf *PlateFans, FansId int) bool {
	log1 := log.NewLog()
	p.db = Init("plate")
	err := p.db.Where("plate_id", pf.PlateId).Find(&pf).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if (FansId-1)/8+1 > len(pf.FansID)/3 {
		return false //不存在
	}
	cpids := pf.FansID[((FansId-1)/8+1)*3-3 : ((FansId-1)/8+1)*3]
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
		pf.FansID = pf.FansID[0:((FansId-1)/8+1)*3-3] + s2 + pf.FansID[((FansId-1)/8+1)*3:]
		pf.FansNumber = pf.FansNumber - 1
		err := p.db.Where("plate_id=?", pf.PlateId).Updates(&pf).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return true
	}
	return false
}

func (p *PlateDB) SelectEasyManage(em *EasyManage, ManageId int) bool {
	log1 := log.NewLog()
	p.db = Init("plate")
	err := p.db.Where("plate_id", em.PlateId).Find(&em).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	if (ManageId-1)/8+1 > len(em.ManageID)/3 {
		return false
	}
	cpids := em.ManageID[((ManageId-1)/8+1)*3-3 : ((ManageId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if 1<<((ManageId-1)%8)&s != 0 {
		return true
	}
	return false
}

func (p *PlateDB) SelectAllEasyManage(em *EasyManage) []int {
	log1 := log.NewLog()
	p.db = Init("plate")
	err := p.db.Where("plate_id", em.PlateId).Find(&em).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	//fmt.Println(cp.ConcernPlatesID)
	easymanage := make([]int, em.ManageNumber)
	sum := 0
	for i := 0; i < len(em.ManageID)/3; i++ {
		s := em.ManageID[3*i : 3*i+3]
		i2, err2 := strconv.Atoi(s)
		if err2 != nil {
			log1.Error(err2.Error(), "")
		}
		for j := 1; j <= 8 && i2 > 0; j++ {
			if i2&1 != 0 {
				easymanage[sum] = 8*i + j
				sum++
			}
			i2 >>= 1
		}
	}
	return easymanage
}

func (p *PlateDB) InsertEasyManage(em *EasyManage, ManageId int) int {
	log1 := log.NewLog()
	p.db = Init("plate")
	err := p.db.Where("plate_id", em.PlateId).Find(&em).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	l := (ManageId-1)/8 + 1
	cpid_len := len(em.ManageID) / 3
	where := 1 << ((ManageId - 1) % 8)
	if l > cpid_len {
		for i := 1; i < l-cpid_len; i++ {
			em.ManageID = em.ManageID + "000"
		}
		s2 := strconv.Itoa(where)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		em.ManageID = em.ManageID + s2
		em.ManageNumber = em.ManageNumber + 1
		err := p.db.Where("plate_id=?", em.PlateId).Updates(&em).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return 2
	} else {
		cpids := em.ManageID[((ManageId-1)/8+1)*3-3 : ((ManageId-1)/8+1)*3]
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
			em.ManageID = em.ManageID[0:((ManageId-1)/8+1)*3-3] + s2 + em.ManageID[((ManageId-1)/8+1)*3:]
			em.ManageNumber = em.ManageNumber + 1
			err1 := p.db.Where("plate_id=?", em.PlateId).Updates(&em).Error
			if err1 != nil {
				log1.Error(err1.Error(), "")
			}
			return 2
		}
	}
}

func (p *PlateDB) DeleteEasyManage(em *EasyManage, ManageId int) bool {
	log1 := log.NewLog()
	p.db = Init("plate")
	err := p.db.Where("plate_id", em.PlateId).Find(&em).Error
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if (ManageId-1)/8+1 > len(em.ManageID)/3 {
		return false //不存在
	}
	cpids := em.ManageID[((ManageId-1)/8+1)*3-3 : ((ManageId-1)/8+1)*3]
	s, err := strconv.Atoi(cpids)
	if err != nil {
		log1.Error(err.Error(), "")
	}
	if 1<<((ManageId-1)%8)&s != 0 {
		s = s & (^(1 << ((ManageId - 1) % 8)))
		//fmt.Println(s)
		s2 := strconv.Itoa(s)
		if len(s2) == 1 {
			s2 = "00" + s2
		} else if len(s2) == 2 {
			s2 = "0" + s2
		}
		em.ManageID = em.ManageID[0:((ManageId-1)/8+1)*3-3] + s2 + em.ManageID[((ManageId-1)/8+1)*3:]
		em.ManageNumber = em.ManageNumber - 1
		err := p.db.Where("plate_id=?", em.PlateId).Updates(&em).Error
		if err != nil {
			log1.Error(err.Error(), "")
		}
		return true
	}
	return false
}
