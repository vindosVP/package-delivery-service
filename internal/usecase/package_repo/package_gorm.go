package package_repo

import (
	"clean-architecture-service/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PackageRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PackageRepo {
	return &PackageRepo{
		db: db,
	}
}

func (p *PackageRepo) GetPackages(ownerID uuid.UUID) ([]entity.Package, error) {
	var packages []entity.Package
	err := p.db.Model(&entity.Package{}).Where("owner_id = ?", ownerID).Find(&packages).Error
	return packages, err
}

func (p *PackageRepo) Create(pack *entity.Package) (*entity.Package, error) {
	return pack, p.db.Create(&pack).Error
}

func (p *PackageRepo) FindByID(packageID uuid.UUID) (*entity.Package, error) {
	var pack entity.Package
	err := p.db.Model(&entity.Package{}).Where("id = ?", packageID).First(&pack).Error
	return &pack, err
}

func (p *PackageRepo) Update(pack *entity.Package) (*entity.Package, error) {
	var newPack entity.Package
	err := p.db.Model(entity.Package{}).Where("id = ?", pack.ID).Updates(pack).First(&newPack).Error
	return &newPack, err
}

func (p *PackageRepo) PackageExistsByID(packageID uuid.UUID) (bool, error) {
	var count int64
	err := p.db.Model(&entity.Package{}).Where("id = ?", packageID).Count(&count).Error
	return count > 0, err
}
