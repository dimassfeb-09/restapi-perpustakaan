-- MySQL dump 10.13  Distrib 8.0.30, for macos12.4 (arm64)
--
-- Host: localhost    Database: library_db
-- ------------------------------------------------------
-- Server version	8.0.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `books` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `category_id` int NOT NULL,
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `stock` int NOT NULL,
  `products_status` enum('out_of_stock','in_stock') DEFAULT NULL,
  `img_url` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_books_categories` (`category_id`),
  CONSTRAINT `fk_books_categories` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (7,'Dimas1',6,'2022-09-11 13:42:12',732234,'in_stock',NULL),(10,'Berlari Mengejar Cita-cita',4,'2022-09-11 13:43:35',1,'in_stock',NULL),(11,'Terbitlah Terang',4,'2022-09-11 13:43:38',1,'in_stock',NULL),(12,'Cinta Itu Terlarang',4,'2022-09-11 13:44:20',1,'in_stock',NULL),(13,'Apa daya',4,'2022-09-11 06:47:21',1,'in_stock',NULL),(14,'Pelajaran Hari Ini 1',4,'2022-09-11 13:48:43',1,'in_stock',NULL),(15,'Pelajaran Hari  2',4,'2022-09-11 13:48:48',1,'in_stock',NULL),(16,'Pelajaran Hari 3',4,'2022-09-11 13:49:44',1,'in_stock',NULL),(17,'Pelajaran Hari 4',4,'2022-09-11 13:49:47',1,'in_stock',NULL),(19,'Pelajaran Hari 5',4,'2022-09-11 13:51:11',1,'in_stock',NULL),(21,'Dimas',4,'2022-09-11 15:15:07',5,'in_stock',NULL),(22,'Dimas',4,'2022-09-11 15:16:05',5,'in_stock',NULL),(26,'Dimas',4,'2022-09-11 15:40:10',1,'in_stock',NULL),(27,'Last add 27',5,'2022-09-11 15:56:10',1,'in_stock',NULL),(28,'Dimas1',6,'2022-09-15 14:54:29',732234,'in_stock',NULL),(29,'Dimas1',6,'2022-09-15 14:56:58',732234,'in_stock',NULL),(30,'Dimas1',6,'2022-09-15 14:57:33',732234,'in_stock',NULL),(31,'Dimaszsas',6,'2022-09-15 15:09:24',10,'in_stock',NULL),(32,'Dimaszsas',6,'2022-09-15 15:10:43',10,'in_stock',NULL),(33,'Dimaszsas',6,'2022-09-15 15:11:51',10,'in_stock',NULL),(34,'Dimaszsas',6,'2022-09-15 15:11:52',10,'in_stock',NULL),(35,'Buku Kita',6,'2022-09-15 22:40:07',10,'in_stock','http://helloworld.jpg'),(36,'Buku Kita',6,'2022-09-15 22:41:11',10,'in_stock','http://helloworld.jpg'),(37,'Buku Kita',6,'2022-09-15 22:41:15',10,'in_stock','http://helloworld.jpg'),(38,'Buku Kita',6,'2022-09-15 22:41:16',10,'in_stock','http://helloworld.jpg'),(39,'Buku Kita',6,'2022-09-15 22:41:19',10,'in_stock','http://helloworld.jpg'),(40,'Buku Kita',6,'2022-09-15 22:41:49',10,'in_stock','http://helloworld.jpg'),(41,'Buku Kita',6,'2022-09-15 22:41:58',10,'in_stock','http://terminal.jpg');
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_uniqe` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (6,'asasdasdd'),(5,'Dimas'),(14,'Dimas1'),(15,'Dimaszsas'),(2,'Movivasi'),(4,'Pecitraan'),(13,'Ucokasd'),(11,'Ucokasdasd');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `guestbook`
--

DROP TABLE IF EXISTS `guestbook`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `guestbook` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `officer_id` int NOT NULL,
  `book_id` int NOT NULL,
  `start_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `end_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_guestbook_users` (`user_id`),
  KEY `fk_guestbook_officers` (`officer_id`),
  KEY `fk_guestbook_books` (`book_id`),
  CONSTRAINT `fk_guestbook_books` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`),
  CONSTRAINT `fk_guestbook_officers` FOREIGN KEY (`officer_id`) REFERENCES `officers` (`id`),
  CONSTRAINT `fk_guestbook_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `guestbook`
--

LOCK TABLES `guestbook` WRITE;
/*!40000 ALTER TABLE `guestbook` DISABLE KEYS */;
INSERT INTO `guestbook` VALUES (18,72,1,10,'2022-09-15 19:19:36','2022-12-05 15:20:20'),(19,74,1,7,'2022-09-15 12:30:21','2022-12-20 15:20:20'),(20,75,1,13,'2022-09-15 12:30:31','2022-12-20 15:20:20'),(21,76,1,15,'2022-09-15 12:30:41','2022-12-20 15:20:20'),(22,77,1,16,'2022-09-15 12:30:42','2022-12-20 15:20:20'),(23,78,1,11,'2022-09-15 12:33:19','2022-12-20 15:20:20'),(24,79,1,7,'2022-09-15 12:44:29','2222-12-05 15:20:20'),(25,80,1,7,'2022-09-15 12:46:53','2022-09-16 15:20:20'),(26,72,1,11,'2022-09-15 12:48:15','2022-09-16 15:20:20'),(31,75,1,14,'2022-09-15 12:30:31','2022-12-20 15:20:20'),(32,75,1,15,'2022-09-15 12:30:31','2022-12-20 15:20:20'),(33,75,1,16,'2022-09-15 12:30:31','2022-12-20 15:20:20'),(34,75,1,17,'2022-09-15 12:30:31','2022-12-20 15:20:20'),(35,72,1,7,'2022-09-15 13:50:37','2022-09-16 15:20:20');
/*!40000 ALTER TABLE `guestbook` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `officers`
--

DROP TABLE IF EXISTS `officers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `officers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `position` varchar(50) NOT NULL,
  `phone` varchar(15) NOT NULL,
  `address` text,
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `officers`
--

LOCK TABLES `officers` WRITE;
/*!40000 ALTER TABLE `officers` DISABLE KEYS */;
INSERT INTO `officers` VALUES (1,'Dimas1','staff','732234','Jalan Jalan 1','2022-09-11 17:43:05'),(3,'Dimas','staff','123649','Jalan Jalan','2022-09-15 14:50:49'),(4,'Dimas','staff','92347','Jalan Jalan','2022-09-15 14:51:16'),(5,'Dimas','staff','92347','Jalan Jalan','2022-09-15 14:51:17'),(6,'Dimaszsas','staff','234789','jalan jalan','2022-09-15 15:12:27');
/*!40000 ALTER TABLE `officers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT NULL,
  `username` varchar(20) DEFAULT NULL,
  `password` varchar(20) DEFAULT NULL,
  `email` varchar(20) DEFAULT NULL,
  `level` varchar(20) DEFAULT NULL,
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username_uniqe` (`username`),
  UNIQUE KEY `email_uniqe` (`email`),
  FULLTEXT KEY `users_search` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (72,'Dimas','qaswasder','qwasde','awed@gmail.com','Mahasiswa','2022-09-10 19:01:30'),(74,'dimas','dimassfeb','Automenu09','qwer1@gmail.com','mahasiswa','2022-09-13 15:32:32'),(75,'Dimas Febriyanto','dimassfeb01','automenu09','dimassfeab@gmail.com','mahasiswa','2022-09-13 15:38:39'),(76,'Dimas Febriyanto','dimassfeb02','automenu09','dsfeab@gmail.com','mahasiswa','2022-09-13 15:39:18'),(77,'asasdasdd','ajksdh','ajksdh','jkhasd','mahasiswa','2022-09-15 14:16:59'),(78,'Dimas','abdjns','kasd','ads','mahasiswa','2022-09-15 14:19:14'),(79,'Dimas','aabdjns','bjkasd','adas','mahasiswa','2022-09-15 14:19:49'),(80,'Dimas','abd098jn','bjkasd','aASdas','mahasiswa','2022-09-15 14:21:17'),(81,'Dimaszsas','dimas123','12324q','qwuye','mahasiswa','2022-09-15 15:11:39');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-16 13:36:32
