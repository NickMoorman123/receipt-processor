package store

import (
	"context"
	"log"
	"os"

	"github.com/NickMoorman123/receipt-processor/errors" 
	"github.com/NickMoorman123/receipt-processor/objects"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger" 
    "github.com/google/uuid"
)

type pg struct {
	db *gorm.DB
}

func NewPostgresReceiptStore(conn string) IReceiptStore {
	db, err := gorm.Open(postgres.Open(conn),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Info,
					Colorful: true,
				},
			),
		},
	)
	if err != nil {
		panic("Enable to connect to database: " + err.Error())
	}
	if err := db.AutoMigrate(&objects.Item{}); err != nil {
		panic("Enable to migrate database: " + err.Error())
	}
	if err := db.AutoMigrate(&objects.Receipt{}); err != nil {
		panic("Enable to migrate database: " + err.Error())
	}
	
	return &pg{db: db}
}

func (p *pg) Get(ctx context.Context, in *objects.GetRequest) (*objects.Receipt, error) {
	receipt := &objects.Receipt{}
	err := p.db.WithContext(ctx).Take(receipt, "uuid = ?", in.UUID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrReceiptNotFound
	}
	return receipt, err
}

func (p *pg) Process(ctx context.Context, in *objects.ProcessRequest) error {
	if in.Receipt == nil {
		return errors.ErrObjectIsRequired
	}
	//in.Receipt.ID = GenerateUniqueID()
	in.Receipt.UUID = uuid.NewString()
	err := p.db.WithContext(ctx).
		Create(in.Receipt).
		Error
	if err != nil {
		return err
	}

	// for _, item := range in.Receipt.Items {
	// 	//item.ID = fmt.Sprint(in.Receipt.ID, index)
	// 	err = p.db.WithContext(ctx).
	// 		Create(item).
	// 		Error
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}