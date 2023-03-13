package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// module import 
	// golang-fiber-postgre-gorm
	"github.com/hzlnqodrey/golang-fiber-postgre-gorm/config"
	"github.com/hzlnqodrey/golang-fiber-postgre-gorm/model"
)