-- --------------------------------------------------------
-- Хост:                         127.0.0.1
-- Версия сервера:               PostgreSQL 12.5 (Ubuntu 12.5-1.pgdg20.04+1) on x86_64-pc-linux-gnu, compiled by gcc (Ubuntu 9.3.0-17ubuntu1~20.04) 9.3.0, 64-bit
-- Операционная система:         
-- HeidiSQL Версия:              11.1.0.6116
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Дамп структуры для таблица public.users
CREATE TABLE IF NOT EXISTS "users" (
	"id" INTEGER NOT NULL DEFAULT 'nextval(''users_id_seq''::regclass)',
	"first_name" VARCHAR(1024) NOT NULL,
	"last_name" VARCHAR(1024) NULL DEFAULT NULL,
	"username" VARCHAR(1024) NULL DEFAULT NULL,
	"language_code" VARCHAR(10) NULL DEFAULT NULL,
	"can_join_groups" BOOLEAN NULL DEFAULT NULL,
	"can_read_all_group_messages" BOOLEAN NULL DEFAULT NULL
);

-- Экспортируемые данные не выделены.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
