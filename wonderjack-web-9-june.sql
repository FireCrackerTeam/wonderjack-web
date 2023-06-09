CREATE DATABASE  IF NOT EXISTS `wonderjack_web` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `wonderjack_web`;
-- MySQL dump 10.13  Distrib 8.0.33, for macos13 (arm64)
--
-- Host: wonderjack-mysql.mysql.database.azure.com    Database: wonderjack_web
-- ------------------------------------------------------
-- Server version	5.7.42-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `available_codes`
--

DROP TABLE IF EXISTS `available_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `available_codes` (
  `code_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` longtext NOT NULL,
  PRIMARY KEY (`code_id`)
) ENGINE=InnoDB AUTO_INCREMENT=301 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `available_codes`
--

LOCK TABLES `available_codes` WRITE;
/*!40000 ALTER TABLE `available_codes` DISABLE KEYS */;
INSERT INTO `available_codes` VALUES (1,'91ba17ed96'),(2,'e57aeb0aa4'),(3,'6c44fe3f6f'),(4,'f8b22b89b3'),(5,'1fe4b445b7'),(6,'329a25d861'),(7,'b0d7ee38dd'),(8,'01e21ea789'),(9,'53dad33e18'),(10,'90ed91cc0f'),(11,'6442cf36b2'),(12,'00ee52b99a'),(13,'174faf8430'),(14,'49b674e961'),(15,'8a0ecdd83f'),(16,'51520b9b6a'),(17,'fc871f3d71'),(18,'403a01d472'),(19,'aa040388c3'),(20,'5108f0bbf0'),(21,'51a0109434'),(22,'e35b33421b'),(23,'0023b2edaa'),(24,'f3464aef7a'),(25,'e51e13b7b9'),(26,'5b7c9d47e5'),(27,'b366ef3057'),(28,'4a0927ac04'),(29,'d35347c8f4'),(30,'e410e79e02'),(31,'18e758c3d9'),(32,'6a5c7e8f74'),(33,'c262a5233a'),(34,'72f4cb3f92'),(35,'36a2931454'),(36,'5cbdc75166'),(37,'f65bb1aec5'),(38,'efa8b0df31'),(39,'993ca29ce2'),(40,'f326d12308'),(41,'1d446d9ea6'),(42,'092e862c8c'),(43,'7a12b27c70'),(44,'053d3f131b'),(45,'2ae4ff9902'),(46,'a1b242e751'),(47,'11c0c8aa13'),(48,'450e502645'),(49,'362bbf9c8a'),(50,'6dad5e8fd1'),(51,'89524c282a'),(52,'fb5310dc29'),(53,'3d98765783'),(54,'e644513ee3'),(55,'6aaeaa49d2'),(56,'8a96457874'),(57,'3616e582f2'),(58,'cf17980c5b'),(59,'64ffb2f2fb'),(60,'6919487ee3'),(61,'6b2818bc5f'),(62,'729a67656b'),(63,'420e354d49'),(64,'e98a2162ed'),(65,'d72fa6e867'),(66,'9df1e783a8'),(67,'b1086fc303'),(68,'491a1d11c6'),(69,'2f5fb51380'),(70,'d91ba26ebd'),(71,'15074683fd'),(72,'6f4636df3f'),(73,'8299634497'),(74,'f2792c366e'),(75,'69ecd3ef13'),(76,'bc80c31a8c'),(77,'56125403de'),(78,'9fc06e3ba9'),(79,'e6b0634ff6'),(80,'72ddb1b81b'),(81,'a2351b9815'),(82,'b1fa0c5eb8'),(83,'188e4d343a'),(84,'dbd69e77ac'),(85,'491b759a27'),(86,'1554502a75'),(87,'f08fa98bfe'),(88,'30ca81c337'),(89,'fd286551a5'),(90,'d3edd9fb7b'),(91,'362f9ecc38'),(92,'55c268c050'),(93,'e81d42fc4a'),(94,'998642930f'),(95,'fa75348e70'),(96,'70b7082ec2'),(97,'2461fe4bd1'),(98,'edc9ec27bd'),(99,'a99a9b1bb4'),(100,'21fe19baf4'),(101,'84b8fa69f2'),(102,'bac571b765'),(103,'9a6ee426c1'),(104,'e28f76b734'),(105,'e9ed743653'),(106,'9411492175'),(107,'c555c697a0'),(108,'ed58c7d43e'),(109,'4e985b5adf'),(110,'d62ac13970'),(111,'e782dfaf9f'),(112,'ec211bc0d7'),(113,'8cecd968e8'),(114,'209e7f7ab3'),(115,'8b15f8a779'),(116,'0d9f07f614'),(117,'8efcc89929'),(118,'7ef9171d38'),(119,'ef480cac20'),(120,'6bdd4c1783'),(121,'a22984031b'),(122,'1ccec68a4d'),(123,'c0964365c7'),(124,'f42cb99647'),(125,'0836f36514'),(126,'d0c9823385'),(127,'a9e3001aab'),(128,'9304038069'),(129,'7080044a91'),(130,'72f52afdf9'),(131,'70e87ce940'),(132,'60341dd1b6'),(133,'109719c86b'),(134,'91c57e67ee'),(135,'9ec2e00060'),(136,'ec72f404ac'),(137,'9273a47c79'),(138,'53a8c9860d'),(139,'31f9ba02b4'),(140,'ae0dafe40b'),(141,'04c0daee4e'),(142,'8467afcf24'),(143,'b18f3bb186'),(144,'aba59e851b'),(145,'3ecfd08ca3'),(146,'ca0004beef'),(147,'f14afcbff9'),(148,'30127af895'),(149,'713b50c65f'),(150,'fd2faa6a18'),(151,'7c3e085ae6'),(152,'0f86433fa4'),(153,'fa5c1bb723'),(154,'4fabf54517'),(155,'008c5690dd'),(156,'3ed4c96cb1'),(157,'219e35be7b'),(158,'26e7cfb22e'),(159,'0edfe4db80'),(160,'369216a87b'),(161,'5b290767b5'),(162,'fd0eb13620'),(163,'2baba6ea90'),(164,'51ed7d5f94'),(165,'14289e57d6'),(166,'c70989898b'),(167,'82dd6c01ca'),(168,'a67224e723'),(169,'9273f25ed6'),(170,'b6aacc1bee'),(171,'47fbed96eb'),(172,'40a604d9c5'),(173,'45a9b0d3ed'),(174,'c79d2cb529'),(175,'7478cea66b'),(176,'da23a6e05b'),(177,'88ff4a7504'),(178,'974228e361'),(179,'9c714f9071'),(180,'bdbd77cdb0'),(181,'f33104c30d'),(182,'fe9b03e9ed'),(183,'ae6a7255fe'),(184,'92c6276962'),(185,'8d5bd28a33'),(186,'f505a0c792'),(187,'01bf11f340'),(188,'7a7483c79c'),(189,'a9984a0914'),(190,'a7444e3f39'),(191,'f16d2b3232'),(192,'755ff5e81f'),(193,'2930593cb0'),(194,'e59e9a44eb'),(195,'59d8e8fc57'),(196,'ac07501275'),(197,'646935cd5b'),(198,'d5f09d41e8'),(199,'99b9e769a7'),(200,'5a7b12deb4'),(201,'df673648f4'),(202,'82119e1944'),(203,'95b1e9894c'),(204,'753252a3c5'),(205,'e5b98cf65e'),(206,'10abb00480'),(207,'452ac36675'),(208,'e669be25ea'),(209,'f24125b3e6'),(210,'3b164f7bb0'),(211,'52e64354b6'),(212,'c45fb6d802'),(213,'affc77e5e9'),(214,'b17a0d65f6'),(215,'46a45a9c58'),(216,'78ea8e66f3'),(217,'b097a0f292'),(218,'72cbc18a50'),(219,'670d3ae54e'),(220,'d4b944b081'),(221,'a424a3393f'),(222,'0b98f1094f'),(223,'dfa98acc7b'),(224,'6d1e50039a'),(225,'a3762c91df'),(226,'7a38b94666'),(227,'72bd9c6b7c'),(228,'a71d433956'),(229,'ca3f299cbc'),(230,'efe96e20d3'),(231,'275c94e064'),(232,'ec77b82e85'),(233,'b1bca8fa5d'),(234,'b1a48b4f9a'),(235,'7c905ffc29'),(236,'a84ac77a0c'),(237,'abc578456a'),(238,'a40afeaa72'),(239,'b13cd98371'),(240,'518f07b2e9'),(241,'21ecb9dffb'),(242,'d6d57acb92'),(243,'1b04da09aa'),(244,'04de3821de'),(245,'5d596ac87a'),(246,'f6a5631692'),(247,'f61c85ed31'),(248,'9042703d8c'),(249,'15ce45b52a'),(250,'a757d1842e'),(251,'807f3a755e'),(252,'04cc40fe06'),(253,'d4d613b97f'),(254,'0546af3959'),(255,'00e5862233'),(256,'4534714b77'),(257,'6c559b49bf'),(258,'e875a56042'),(259,'90fa971ee1'),(260,'4786c39f05'),(261,'b04801106f'),(262,'30930e8c0d'),(263,'2116f914c7'),(264,'86d9b94f07'),(265,'e8b4852347'),(266,'15af2f6e85'),(267,'468bd0a448'),(268,'9642a9782d'),(269,'22a8dd1be0'),(270,'c37221ebb2'),(271,'72a1a73785'),(272,'790ab4e6ce'),(273,'b5a1d169d9'),(274,'dbe59a5a67'),(275,'adab189240'),(276,'8035bd13a8'),(277,'2060748a57'),(278,'09ad74d80c'),(279,'fc4d1d0352'),(280,'ec7379bfb7'),(281,'08c96fd459'),(282,'54fe3530c2'),(283,'bfdad60332'),(284,'c79921275d'),(285,'220363454d'),(286,'beeac9c399'),(287,'54cd339330'),(288,'a7b28700f4'),(289,'9bb0669f5e'),(290,'8863b1ee83'),(291,'e50f39667e'),(292,'8d8e446bba'),(293,'54855bf266'),(294,'2c6e5925d2'),(295,'8bc2deb50a'),(296,'3eb714838e'),(297,'3dd889e6a6'),(298,'96f15769f7'),(299,'29a6a5f6a5'),(300,'d21b1a6824');
/*!40000 ALTER TABLE `available_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `redemption_codes`
--

DROP TABLE IF EXISTS `redemption_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `redemption_codes` (
  `redemption_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `code` longtext NOT NULL,
  PRIMARY KEY (`redemption_id`),
  KEY `fk_users_redemption_codes` (`user_id`),
  CONSTRAINT `fk_users_redemption_codes` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `redemption_codes`
--

LOCK TABLES `redemption_codes` WRITE;
/*!40000 ALTER TABLE `redemption_codes` DISABLE KEYS */;
INSERT INTO `redemption_codes` VALUES (8,1,'d21b1a6824'),(15,3,'3dd889e6a6'),(16,4,'91ba17ed96'),(17,6,'36a2931454'),(18,10,'72f4cb3f92'),(19,11,'f8b22b89b3'),(20,12,'fc871f3d71');
/*!40000 ALTER TABLE `redemption_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_name` longtext NOT NULL,
  `user_email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_email` (`user_email`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Hafshy','test1@email3.com','$2a$10$CdWG33nyHh2ZnFZ1OB1wX.lD9mPZ0hJz8IOnOiPmZAmM3uYSyHrOK'),(2,'Wahed','wahed@gmail.com','$2a$10$WG5tqngUGuNOoxcGQ9.nh.oATLGDMv7Oi.JB9/nYn.WH5qltIpmm.'),(3,'WonderJack Tester','wonderjack.game@gmail.com','$2a$10$fydr.e/JuMFhvaB1nV3GXuWf4WjYuvEcOpJNAqRg0xLxL1ifobfAu'),(4,'Lulu','Satyaningm99@gmail.com','$2a$10$Wmmb7YPTpXWrt5ujgM5MTOR25YLLRHfwz3.P/6GZP0TPgSaPufXVu'),(5,'Timmy','nathannorth2005@gmail.com','$2a$10$ELFWrj2oLiqkHThutkbfueXi7lTtO/wXTAynfl08arhklo0tNfG.W'),(6,'Kevin','chrisk.setiawan@gmail.com','$2a$10$HEUjMO.HAw1zwAVHPZj6O.oYMFX1Ac/96eCLqIKfCsya.w/usgwnu'),(7,'Hafshy Yazid Albisthami','hafshy@wonderjackgame.com','$2a$10$aSEVGd6AQVmJ/1NSD0vP/ukDLTVPSR.f7n8OHsQ5/R1qKl0/dAG9u'),(10,'Ranee','Contactsatyaningm@gmail.com','$2a$10$kMdyQWdzQyI.GCa0..YNkeEe0aHW5tK1YvZPBjQrXthFzdN8HfNxK'),(11,'Shirleen','shirleen.goenawan@gmail.com','$2a$10$qvFqhnMO1/orjwlRpyF77OMC0.apPgC5IiLg5A6RRVHY7Y2aqpa/S'),(12,'Catherine','1234kathrina@gmail.com','$2a$10$.wd9RKPL15xsjwE8Ub9T3.Fxvhd/CS4Fizynz376i0MN8VzLmIisC');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'wonderjack_web'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-09 16:42:17
