-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: db_playup
-- ------------------------------------------------------
-- Server version	8.0.20

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
-- Table structure for table `badminton_single_player_list`
--

DROP TABLE IF EXISTS `badminton_single_player_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `badminton_single_player_list` (
  `player_list_id` varchar(100) NOT NULL,
  `player_list_match_id` varchar(100) DEFAULT NULL,
  `player_one` varchar(100) DEFAULT NULL,
  `player_two` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`player_list_id`),
  KEY `player_list_match_id` (`player_list_match_id`),
  CONSTRAINT `badminton_single_player_list_ibfk_1` FOREIGN KEY (`player_list_match_id`) REFERENCES `master_match` (`match_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `badminton_single_player_list`
--

LOCK TABLES `badminton_single_player_list` WRITE;
/*!40000 ALTER TABLE `badminton_single_player_list` DISABLE KEYS */;
INSERT INTO `badminton_single_player_list` VALUES ('00245e66-fe7b-b4e7-8f4f-e5f3af5e22e2','00eb1682-a548-0939-14c0-aaefb6e552eb','user_three','user_one'),('01665820-7526-6e26-7d71-d84a5dbcc3da','ec093a24-e06d-16d8-b08d-5fe726d9d1f5','user_one','user_two'),('142ff9f3-c6fa-f33f-6211-d1d8797e81c4','3ebb9d6a-f784-6f59-3a4c-761a4b7cdff0','user_one','user_two'),('1d0bd4fc-1c30-faa1-8ffd-5355f71d5cb1','9fd61c90-edc6-67cf-b96c-b5e915d875e7','user_suda-diganti','user_four'),('1ecf7fad-bdc8-3c4f-8084-1f7318821865','d36d9425-0d30-85f2-473c-1e0106d29c19','user_one','user_two'),('32561374-a63f-bf34-9857-1e55cd37a59c','2661b035-24b4-1c67-3eae-3beafd186756','user_one','user_two'),('39fcf661-21fc-8ed0-c1da-069b2e9fa52f','27c3a417-38a5-1a44-4b0d-0090862f7fc2','user_four','user_suda-diganti'),('48057c51-7e56-bb6a-247e-8525a0b9ae27','029884f7-19ce-1e8a-8e1c-9a8236a02834','user_two','user_one'),('48ca85e9-a09b-3dfc-0f2c-eee55d4d1007','ec093a24-e06d-16d8-b08d-5fe726d9d1f5','user_one','user_two'),('4b652923-8872-9e84-b761-535faddaadd8','d36d9425-0d30-85f2-473c-1e0106d29c19','user_one','user_two'),('4cd0d4e0-f91b-aad6-1357-91aa84c2cafb','1031aea4-96b3-5f5a-ad36-7e86637d0e2f','user_suda-diganti','user_four'),('4ee0114c-efa2-452c-a79d-db1131357c02','d36d9425-0d30-85f2-473c-1e0106d29c19','user_one','user_two'),('55a7b54c-e32e-9ba3-b50b-86908f8dab7a','ini ngaco','user_one','user_two'),('5e3f16f3-adc2-5c08-2f23-c3946a203330','45e34f60-a7d1-7ca6-36de-218a7f2fff8c','user_one','user_two'),('7b446616-db71-b7f6-93bf-1bbe5f35b50a','3ebb9d6a-f784-6f59-3a4c-761a4b7cdff0','user_one','user_two'),('a644495a-2f17-cae3-7a9d-1e226ebb19a0','00eb1682-a548-0939-14c0-aaefb6e552eb','user_three','user_one'),('a8aef3cd-420d-8744-890a-665a673fda12','d36d9425-0d30-85f2-473c-1e0106d29c19','user_one','user_two'),('af32489a-10bd-2948-67d8-039569199691','33eae5b5-35a1-a229-5f72-8e7452a2f5e5','user_one','user_two'),('bebabfc0-f2b6-7aeb-2edf-a4e591e7e85c','33eae5b5-35a1-a229-5f72-8e7452a2f5e5','user_one','user_two'),('c259d1da-d496-69b7-a1b1-813b53028a94','26f58bbc-26df-9797-2cb4-715304a1287c','user_one','user_four'),('d0598e62-d06e-2ee3-640b-af1b0f2b400c','2661b035-24b4-1c67-3eae-3beafd186756','user_one','user_two'),('f452491d-b446-98fa-6f84-bde4d6e1bba2','5e948ae8-b1c7-b1ac-1093-bde90df963c9','user_one_dev','user_two_dev'),('ff064c43-ad03-95b3-b7a0-eca8af362268','029884f7-19ce-1e8a-8e1c-9a8236a02834','user_two','user_one');
/*!40000 ALTER TABLE `badminton_single_player_list` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master_chat`
--

DROP TABLE IF EXISTS `master_chat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `master_chat` (
  `chat_id` varchar(50) DEFAULT NULL,
  `chat_match_id` varchar(100) DEFAULT NULL,
  `chat_sender_id` varchar(50) DEFAULT NULL,
  `chat_receiver_id` varchar(50) DEFAULT NULL,
  `chat_message` longtext,
  KEY `chat_match_id` (`chat_match_id`),
  CONSTRAINT `master_chat_ibfk_1` FOREIGN KEY (`chat_match_id`) REFERENCES `master_match` (`match_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master_chat`
--

LOCK TABLES `master_chat` WRITE;
/*!40000 ALTER TABLE `master_chat` DISABLE KEYS */;
INSERT INTO `master_chat` VALUES ('0038fa43-2eff-b776-ffd1-b156875d9c35','d36d9425-0d30-85f2-473c-1e0106d29c19','20d4574b-b5ee-8df6-5c05-184f1e29d493','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','ini dari user_pertama'),('f465857a-c6db-2fb4-d490-dd8001673bcb','d36d9425-0d30-85f2-473c-1e0106d29c19','20d4574b-b5ee-8df6-5c05-184f1e29d493','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','ini dari user_pertama'),('e1651123-4fd2-03fd-ffbd-3d04cdfc1b61','00eb1682-a548-0939-14c0-aaefb6e552eb','9c267c5c-52bd-a55e-5d88-cc53a70128c8','9c267c5c-52bd-a55e-5d88-cc53a70128c8','hello'),('1dab0ba2-2cb2-2750-9fea-0fe5e27e1234','00eb1682-a548-0939-14c0-aaefb6e552eb','20d4574b-b5ee-8df6-5c05-184f1e29d493','20d4574b-b5ee-8df6-5c05-184f1e29d493','Humana no ?'),('603acfc5-a9d0-521a-f4ed-01fb1bcea431','00eb1682-a548-0939-14c0-aaefb6e552eb','20d4574b-b5ee-8df6-5c05-184f1e29d493','20d4574b-b5ee-8df6-5c05-184f1e29d493','Yang booking soaps ?'),('7ddfd0e7-ef6a-b791-0ad3-67e9b4114170','00eb1682-a548-0939-14c0-aaefb6e552eb','9c267c5c-52bd-a55e-5d88-cc53a70128c8','9c267c5c-52bd-a55e-5d88-cc53a70128c8','yaudah aku aja'),('c2f4d576-afc9-7970-87da-a5d4e40c827a','2661b035-24b4-1c67-3eae-3beafd186756','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','Humana ?\nHumana ?'),('6678ca26-7230-94fe-70e5-73a996aaf8fd','2661b035-24b4-1c67-3eae-3beafd186756','20d4574b-b5ee-8df6-5c05-184f1e29d493','20d4574b-b5ee-8df6-5c05-184f1e29d493','yaudah SKU AKA yg atur jadwalnya '),('46a071e3-d7a4-45c8-160c-e5590f825623','2661b035-24b4-1c67-3eae-3beafd186756','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','oke'),('ee5f6b18-3dcb-b78f-aeec-804c177163bc','33eae5b5-35a1-a229-5f72-8e7452a2f5e5','20d4574b-b5ee-8df6-5c05-184f1e29d493','20d4574b-b5ee-8df6-5c05-184f1e29d493','tak atur yo '),('092ab03c-df3d-79c6-bc59-9d54cfef9b88','029884f7-19ce-1e8a-8e1c-9a8236a02834','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','awdaw'),('a086f68d-ee39-4b2d-4f65-22792cde6db6','029884f7-19ce-1e8a-8e1c-9a8236a02834','20d4574b-b5ee-8df6-5c05-184f1e29d493','20d4574b-b5ee-8df6-5c05-184f1e29d493','2312131231');
/*!40000 ALTER TABLE `master_chat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master_match`
--

DROP TABLE IF EXISTS `master_match`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `master_match` (
  `match_id` varchar(50) NOT NULL,
  `match_sport_id` int DEFAULT NULL,
  PRIMARY KEY (`match_id`),
  KEY `match_sport_id` (`match_sport_id`),
  CONSTRAINT `master_match_ibfk_1` FOREIGN KEY (`match_sport_id`) REFERENCES `master_sport` (`sport_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master_match`
--

LOCK TABLES `master_match` WRITE;
/*!40000 ALTER TABLE `master_match` DISABLE KEYS */;
INSERT INTO `master_match` VALUES ('00eb1682-a548-0939-14c0-aaefb6e552eb',1),('029884f7-19ce-1e8a-8e1c-9a8236a02834',1),('2661b035-24b4-1c67-3eae-3beafd186756',1),('26f58bbc-26df-9797-2cb4-715304a1287c',1),('33eae5b5-35a1-a229-5f72-8e7452a2f5e5',1),('3ebb9d6a-f784-6f59-3a4c-761a4b7cdff0',1),('5e948ae8-b1c7-b1ac-1093-bde90df963c9',1),('ec093a24-e06d-16d8-b08d-5fe726d9d1f5',1);
/*!40000 ALTER TABLE `master_match` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master_rank`
--

DROP TABLE IF EXISTS `master_rank`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `master_rank` (
  `rank_id` varchar(50) NOT NULL,
  `rank_user_id` varchar(50) DEFAULT NULL,
  `rank_user_match_count` int DEFAULT NULL,
  `rank_user_grade_count` int DEFAULT NULL,
  PRIMARY KEY (`rank_id`),
  KEY `rank_user_id` (`rank_user_id`),
  CONSTRAINT `master_rank_ibfk_1` FOREIGN KEY (`rank_user_id`) REFERENCES `master_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master_rank`
--

LOCK TABLES `master_rank` WRITE;
/*!40000 ALTER TABLE `master_rank` DISABLE KEYS */;
INSERT INTO `master_rank` VALUES ('381c0ebc-8d0d-99f2-197f-cf7cb9d9359f','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9',2,1),('46eeff29-3775-c3ea-f2ab-1071c564ccf1','c031e52d-11f0-565d-237d-cf9a7497cd36',0,0),('68d30789-1ef7-75e8-8a0b-ec1c3733fbd2','20d4574b-b5ee-8df6-5c05-184f1e29d493',3,2),('bb57a528-e14d-e066-3306-60c4618b5ef4','ad208b3d-5159-3ad2-eccf-8da872bae25e',0,0),('cad8ec67-bafd-76ad-27b5-0cf48c14ee5a','9c267c5c-52bd-a55e-5d88-cc53a70128c8',6,5),('d697f725-9e89-7d4f-8947-9d84f8d0084b','296bf0dc-a0b6-9b35-9152-880296168dcf',0,0);
/*!40000 ALTER TABLE `master_rank` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master_result`
--

DROP TABLE IF EXISTS `master_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `master_result` (
  `result_id` varchar(100) NOT NULL,
  `result_match_id` varchar(100) DEFAULT NULL,
  `result_winner` varchar(100) DEFAULT NULL,
  `result_loser` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`result_id`),
  KEY `result_match_id` (`result_match_id`),
  CONSTRAINT `master_result_ibfk_1` FOREIGN KEY (`result_match_id`) REFERENCES `master_match` (`match_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master_result`
--

LOCK TABLES `master_result` WRITE;
/*!40000 ALTER TABLE `master_result` DISABLE KEYS */;
INSERT INTO `master_result` VALUES ('02bc06f5-ec8e-cb22-8c09-466cfde73893','64b261ad-b63c-816d-d1e2-10bc2a8d6f7e','unknown','unknown'),('030c7d15-9813-4a29-3707-14f040163f3c','55273e58-246c-79a6-66a8-fef00d5e7acc','unknown','unknown'),('0732d264-76c9-5e59-dcca-8efa96367109','2af7b6c0-0939-85e9-3586-972ae06a9486','unknown','unknown'),('07afeb26-ca1d-dcf2-e4fa-42f696a2cae3','7c2eeba5-1bfa-e4db-2ce1-cc550ea94ae5','unknown','unknown'),('0bf89001-0fee-629f-759f-a6695c6a6b7b','fd905576-4391-e2ec-12df-34eb708e3378','unknown','unknown'),('0cee5ff1-38d6-43b9-b57a-744016d4d8d8','ddb8f6a3-a291-5ddf-7530-2f0c2936bc9e','unknown','unknown'),('0f16feea-a931-3dec-b6e5-99915d632703','c5b53915-2e9c-dfbf-09da-4465585fad52','unknown','unknown'),('10801d35-b26c-bb1d-0f27-e83049650789','c97ccc62-e5aa-2287-509c-4957737646d8','unknown','unknown'),('12080aca-d32b-f889-b2e8-2344b4587bfe','80937296-f909-1541-65fb-d2da00386c1c','unknown','unknown'),('15d6a615-7595-fbc1-e766-27a61182a988','fd23ba8e-6019-62f7-d2cb-b20c2a166bcb','unknown','unknown'),('160cb224-fbb3-269f-14a1-dbf9b8b279ba','9fd61c90-edc6-67cf-b96c-b5e915d875e7','unknown','unknown'),('1d2bc1f8-acd9-a579-bea7-438433d57e9b','743ff0bf-e7eb-189a-a1d7-cffbd7c13619','unknown','unknown'),('1dd7d3d1-17de-68dc-570b-f889cb779ad0','e56f7021-d5fd-0393-bcff-68b00421620c','unknown','unknown'),('1fc16a2a-75c5-3aa2-f80b-10947f022066','be26f64f-5bf9-0f45-9807-11d68847e406','unknown','unknown'),('26ebd98d-7673-8582-3309-dc8f3bb9128f','472d1988-e47b-f9a9-dfd7-1e0458b063e9','unknown','unknown'),('26ef28e6-4ec7-7bef-1ef0-bc8c5ed41705','2e901cac-eeee-8180-1923-125b59093f1c','unknown','unknown'),('27dce696-7a48-5a9d-1cfe-149cb2d31460','ed43a7c0-0507-4b73-2239-148d6e5e1fe6','unknown','unknown'),('299e8502-887e-0c9b-df3a-9525473f36a6','e7579fe3-e405-fedb-dc35-ae6753f88fbc','unknown','unknown'),('2a85e302-f9b3-5541-8de9-78d3e0c66c02','d21bdf60-f655-1d5f-2f29-eb2eadc9bbaa','unknown','unknown'),('2ce10119-96df-8b17-0854-d68f0ec22e26','319f1d1e-77c2-0868-5a6c-a6fb8eddafda','unknown','unknown'),('2d43def2-c385-99e7-a181-af8b63c1873f','86a48a1b-c661-ac2b-20ec-29cd3cb0b48e','unknown','unknown'),('2d91011f-9f1c-e8b1-b5ef-aa0c950cbe56','bae203bc-5224-3083-35ce-28cfbab6522a','unknown','unknown'),('2f724b2c-6f67-0dd4-ecfc-92fc773e0335','cd3d5779-32ae-12f3-e433-afd755568d27','unknown','unknown'),('31234aa8-cf70-76c9-4f17-81cd56e2d807','6763d04c-0d36-3ca6-141e-d0f3ec494bd8','unknown','unknown'),('34ceb546-e753-d7f1-1276-d84b3b245ed4','1a0517cc-ecd6-a7e3-84ca-bd6809edeb50','unknown','unknown'),('37bc734f-fffe-9cb4-514f-2e0b7f457e51','45e34f60-a7d1-7ca6-36de-218a7f2fff8c','unknown','unknown'),('3c1e6e71-a708-32fc-e1cc-4e016d13f018','660792af-a648-8bc1-e305-5c6ca88a560e','unknown','unknown'),('41173aa9-bd8a-8f45-a5cf-57bc88701ea1','4387f631-b001-d162-a6ee-61b3050c2790','unknown','unknown'),('45327da3-f5ba-fa89-d7be-e9effea407a4','ede47562-0248-e570-0036-7d1a5bc65f30','unknown','unknown'),('45e8fc2f-9071-aac6-518e-adca323f9d90','3042d039-56f8-47a4-dc98-f3751c81af50','unknown','unknown'),('4769094e-452a-23e6-8ce1-6d1a6fb93b3e','52da6488-af40-bc71-9e95-96b3f24dc1b7','unknown','unknown'),('48413f61-8752-5413-a970-5961308f1bff','89f9fa52-6599-11e3-20d3-71b36303f94c','unknown','unknown'),('48f3ddff-a731-d219-5ebc-b23eaff15b50','dbfae8f4-a8b5-2ba8-2a6f-147135eb6615','unknown','unknown'),('4b9709ea-e4dc-370c-c43b-548ca41a1001','d36d9425-0d30-85f2-473c-1e0106d29c19','user_one','user_two'),('4bd6b292-8f48-7710-0472-5f546a08cea6','553bc3dd-af2f-012e-25fe-f84ce6beb2e6','unknown','unknown'),('4c347033-5644-8e6a-4ac7-14a96bb4eed0','583f290b-b281-8f7d-39b6-d559eeb178dc','unknown','unknown'),('4cafd033-f707-7559-54e9-b71eba889259','33eae5b5-35a1-a229-5f72-8e7452a2f5e5','user_one','user_two'),('50d65e5f-ca28-1a09-ebe1-0df28ca2cffb','7c28c56e-e5d3-71a5-d155-81ba45411b60','unknown','unknown'),('536881d2-664c-7a69-4957-590200a32d80','f6efc4b3-f168-3640-5a5b-52a645cce4b7','unknown','unknown'),('54dac114-6f91-2ed2-97f5-f524689dccb9','5dca65e9-459b-00b3-9867-0c2a29a88216','unknown','unknown'),('5667d73c-d213-cbbf-c7dc-4b865156f04d','c426fdb7-25f3-fd0e-97ce-97254ff9d535','unknown','unknown'),('590ce082-a31b-cfb7-ba96-e6148775413e','190d2994-925a-6e70-d015-ab1d22c0fabf','unknown','unknown'),('5bba799c-cf3f-4cf8-752f-8d6386c0a0da','e36ff667-9564-78df-21d2-19162d5fc5dd','unknown','unknown'),('5d3ef63b-3af7-ada0-ad25-cb1e94fea155','ae7310b2-2678-5c56-d5ed-d5515032351a','unknown','unknown'),('5dd3053b-3223-3b4f-3039-0775c81df164','d1125134-893c-db35-fab0-dd2da669bc09','unknown','unknown'),('600da64c-461e-3658-926d-5b1a0878e51d','5e948ae8-b1c7-b1ac-1093-bde90df963c9','unknown','unknown'),('64eee32a-f827-4cda-31cf-b1a37e226c6b','dcc400cc-db85-15b7-3058-15d41771861e','unknown','unknown'),('6590c9b0-a310-6ea1-e523-00b7495f19ff','a17dc55e-f70d-15ca-e2b6-b371b6069312','unknown','unknown'),('66a9fbba-3fc7-40b4-e017-819922c8cc3f','a3e5886a-c183-8db1-f4a2-6b56c45583fb','unknown','unknown'),('69eb8978-0989-aef2-c82d-526160a6e153','07eb374b-8248-05e1-2bce-02cde14c7bb6','unknown','unknown'),('6a1675e3-548c-e55a-e122-bf8ec2f16e7a','46253b77-427f-6980-354d-c0ef8c164621','unknown','unknown'),('6aa0b410-7506-ebca-6f1d-4ce6c31a20c6','558ee24a-fc73-1e5e-1013-9c828e9014c4','unknown','unknown'),('7017fd42-3cea-5585-bab6-379e8d58503b','c701cf20-991b-63dc-5c4d-cfbc9284b1ed','unknown','unknown'),('73be629a-5e21-f31f-978d-a330b2b72f28','f5ab6461-fdc5-c75c-5eb7-f8396b66c7fd','unknown','unknown'),('7756055f-b40c-ac30-93bb-c182aa5d2fae','637c4948-a016-d174-05d7-cab491e1e1f8','unknown','unknown'),('79b8f454-8086-4c2b-38b9-053968d78ecf','e0046db7-0085-f16f-d636-f38ee2157dca','unknown','unknown'),('7f1afd3c-61bf-4b2e-3e7a-1b938d5ee6c1','dc771118-8213-21f4-cd1f-930cb420d49f','unknown','unknown'),('7f4a0bf1-3a3c-742c-9941-11698a07bb24','3ebb9d6a-f784-6f59-3a4c-761a4b7cdff0','user_one','user_two'),('8415ee1e-589d-d872-25c9-b47b7d9cf916','93a94974-dfa9-9ac1-719a-fdb70f110cb6','unknown','unknown'),('8445666c-0f5c-eb49-cc47-09607d4d34f1','7f189eaa-fc75-dd0e-d7a3-5f5fd2f5a08a','unknown','unknown'),('86b9a658-42d7-9f42-95f3-6f72114eb1e9','fed182ae-c91b-63f4-9ef6-e4366ed4431b','unknown','unknown'),('8f4591eb-15ff-816d-70e9-fd9eba52156f','ec093a24-e06d-16d8-b08d-5fe726d9d1f5','user_two','user_one'),('92c3c862-9f33-edbc-bd75-4703ae5a8017','bdb770af-3f3c-5566-573d-632dd965a1c8','unknown','unknown'),('9381f2ec-081a-fa3a-71c9-642e0a5bbe52','1694df92-3951-92bc-a971-4a592f6c7fc9','unknown','unknown'),('9527f37c-8a62-1dc5-2127-4b08f58fb33c','ea12c6ad-b734-2e2c-c9f7-5b42e8027e84','unknown','unknown'),('973c82f8-f72a-80f4-762c-ab7b1f94acfe','27c3a417-38a5-1a44-4b0d-0090862f7fc2','unknown','unknown'),('9ad81ea5-4273-68dd-daf2-53c6ec4dde5c','845f5a72-6586-6ade-cdc3-789fd3110bb0','unknown','unknown'),('9d95ca97-9009-682b-1944-8420f4f2b8fe','180ab85d-88b3-0b9d-81ec-24a54142917e','unknown','unknown'),('a2e9762a-6338-d46d-fd90-955bd0aefd87','48e5d8e0-5d85-8822-3eb5-370b042a7e41','unknown','unknown'),('a78c2420-7814-3242-793c-e0a918102e28','029884f7-19ce-1e8a-8e1c-9a8236a02834','unknown','unknown'),('a7e63dec-92d3-ebf8-8c14-c2b24aa0d49c','47da436c-ce1c-c906-d6d1-9e8bce45c039','unknown','unknown'),('a922058f-ef23-ac87-4e49-519b5f1971c0','ec652dea-716e-74c9-eed8-3243ac166970','unknown','unknown'),('b17d6b98-0c3e-a737-4462-f2d06dc569a7','00eb1682-a548-0939-14c0-aaefb6e552eb','user_one','user_three'),('b1c26f8b-4889-17ad-c2ff-1c2be0988710','0d66205c-671c-da87-8aba-22253c459139','unknown','unknown'),('b240485a-1f62-d839-a802-be060797d1c0','a6783b05-2c15-1b89-9a79-27fa45a9230d','unknown','unknown'),('b2e52970-2d55-a70c-c237-ecf77db7f933','1e1054ad-02ea-a235-4743-3254f0fdb13d','unknown','unknown'),('bbebe1a7-24f8-8745-054a-f8977f62be9e','5f3ba3bb-4706-cab6-c64d-b21e47113895','unknown','unknown'),('bed5794a-ce1d-9e2f-b788-2b16ffd9a98c','33d16412-eb1b-e1b4-da2c-9dcc83c956df','unknown','unknown'),('c1d7c184-7a72-8bfe-5cee-f88a3d29ed10','1d85ff1f-7648-1aa8-58a6-4c378bda078d','unknown','unknown'),('c456e4a8-25d9-b878-c312-80c05652c9d8','8e1e0c5f-b161-7376-738f-fb8b497fb9a1','unknown','unknown'),('c68ca24a-690d-efd6-d62f-661eb31a7876','586252a0-9cd5-bf73-86b6-6f7092ac830a','unknown','unknown'),('c7da9f77-73af-98b2-6648-56a7cf10067d','3c28bfcd-aecf-2f1d-950e-c72cc4ae4c88','unknown','unknown'),('c8234128-defc-7afc-6949-60aec211f177','2661b035-24b4-1c67-3eae-3beafd186756','unknown','unknown'),('cabca698-cd35-705c-52d8-c62a9e915b63','0ddf44e5-a0ee-a183-571e-de50f581c0c9','unknown','unknown'),('cdcb36cc-b3a1-41e9-8741-76a9d9759d22','1031aea4-96b3-5f5a-ad36-7e86637d0e2f','unknown','unknown'),('cf7e1145-3095-7de5-3f23-43ddf650aa15','1d5acd2a-2dac-06a2-7bd9-1adb419d2932','unknown','unknown'),('d28333f9-354b-5eb9-50a1-0935560c8601','0f410171-ed66-cf64-4926-4d279a478de7','unknown','unknown'),('d5f676ca-9526-a970-28fb-eccb26dd04b7','34c6239e-6c1d-baf9-4817-77e4d1755f72','unknown','unknown'),('d77431a6-43e7-3c8d-a73f-386748192876','cf401250-9282-531b-3266-b9fcde28fbf0','unknown','unknown'),('d95831ec-34bb-4526-91de-a9d57066975e','b07a120e-4e6a-0361-4686-d2246d8fcc08','unknown','unknown'),('d974213b-5194-ee38-d485-62b4a8a4f1c0','64765d79-1912-2d69-6b47-0f64946246f7','unknown','unknown'),('da2ce62c-6f4b-1fd9-075a-907872f04ca6','6e213bbb-d40a-3b24-8ef1-b9bf00c0b8cf','unknown','unknown'),('da6195c9-a98a-8585-2df3-76e9c42c4434','d0b3540d-f785-490e-1ff8-40775e4652ba','unknown','unknown'),('dbbcda30-36e7-0437-e35c-043c06d22d46','162b840d-c6c3-0b1a-c488-f6273f70a62b','unknown','unknown'),('dd6535cc-9142-374d-7c7c-c59b2a92aeb3','1bd0fe0b-fa4f-9fcb-3879-4604646abc01','unknown','unknown'),('e0b610cb-f192-3ca9-9db7-6c65556d3804','5cf852cb-6060-edf2-d1af-8661a535229e','unknown','unknown'),('e3d34594-9ae1-4d81-fdd7-bbf9dbeb1856','b478250e-c5dc-e367-acd0-7f7aa127e838','unknown','unknown'),('eb1e236a-f744-3689-4ff4-2f29704456f8','cfbf1dc8-f790-56f8-18ee-2606e572ee33','unknown','unknown'),('f1dfaa60-3603-0116-9497-14756284aeef','26f58bbc-26df-9797-2cb4-715304a1287c','unknown','unknown'),('f28d7992-d8f8-2f84-06ef-bf32f0367e8a','543e6a21-dbdf-4915-ef2d-a72d4ed3ae95','unknown','unknown'),('f2d64f55-bff1-7327-834a-a2360c098cca','762ddb30-32a7-2a3b-c6ef-9cb8842d15dc','unknown','unknown'),('f8f0c73c-fbae-bddc-dc99-7be15d4f542b','a72eb5d1-fe4f-0955-d61f-91dd9a240cd8','unknown','unknown'),('f910186d-258b-16cd-e79f-36cc09bbcf9b','aea8bfd6-54de-243d-f13c-f0b11497fb69','unknown','unknown'),('fb31eb41-090c-c7e6-2a77-984ecd94db2f','ee2b4f64-0587-1716-3432-db97836e2b1c','unknown','unknown'),('fc652d81-b1b6-a4c8-0d66-cbd7955ae65d','7d6af054-caf1-968d-8471-7b7b97433e7e','unknown','unknown'),('fdc133be-1cfb-09d4-3bb7-89ac640e5b40','53a6cce8-801d-adf9-bc89-7c39c03757dc','unknown','unknown'),('fddb8b08-f482-6dac-fe2a-98e58d55d6a6','79ae94f4-6161-fc7c-2374-ac52c7cb3140','unknown','unknown'),('ff4c46de-12b2-db7f-645e-4d898ca721b7','992e4dd2-647b-1c6a-1006-179f75e3e8a5','unknown','unknown');
/*!40000 ALTER TABLE `master_result` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master_schedule`
--

DROP TABLE IF EXISTS `master_schedule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `master_schedule` (
  `schedule_id` varchar(50) NOT NULL,
  `schedule_match_id` varchar(50) DEFAULT NULL,
  `schedule_user_id` varchar(100) DEFAULT NULL,
  `schedule_user_name` varchar(100) DEFAULT NULL,
  `schedele_location` longtext,
  `schedule_time` varchar(50) DEFAULT NULL,
  `schedule_status` varchar(1) DEFAULT NULL,
  `schedule_result` varchar(1) DEFAULT NULL,
  `schedule_opponent` varchar(50) DEFAULT NULL,
  `schedule_opponent_id` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`schedule_id`),
  KEY `schedule_match_id` (`schedule_match_id`),
  KEY `schedule_user_id` (`schedule_user_id`),
  CONSTRAINT `master_schedule_ibfk_1` FOREIGN KEY (`schedule_match_id`) REFERENCES `master_match` (`match_id`),
  CONSTRAINT `master_schedule_ibfk_2` FOREIGN KEY (`schedule_user_id`) REFERENCES `master_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master_schedule`
--

LOCK TABLES `master_schedule` WRITE;
/*!40000 ALTER TABLE `master_schedule` DISABLE KEYS */;
INSERT INTO `master_schedule` VALUES ('07397e81-0a27-3692-1adb-6195e0af8785','ec093a24-e06d-16d8-b08d-5fe726d9d1f5','20d4574b-b5ee-8df6-5c05-184f1e29d493','user_one','SEMARANG, ISTORA SENAYAN','SENIN, 20 DESEMBER 2020','I','L','user_two','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9'),('175742bd-356a-8477-291b-1cf38bcd09bf','2661b035-24b4-1c67-3eae-3beafd186756','20d4574b-b5ee-8df6-5c05-184f1e29d493','user_one','template satu','2020 Oktoberfest 15.00','A','U','user_two','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9'),('6b9fc659-0b68-e551-d52e-b4fea25479cb','33eae5b5-35a1-a229-5f72-8e7452a2f5e5','20d4574b-b5ee-8df6-5c05-184f1e29d493','user_one','lapangan badminton hore','Santi, 10 Oktoberfest 2020 ','I','W','user_two','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9'),('80c9e49e-f944-00fc-7734-1b0df4480cb7','00eb1682-a548-0939-14c0-aaefb6e552eb','9c267c5c-52bd-a55e-5d88-cc53a70128c8','user_three','template pertama','18 Oktoberfest 2020','I','L','user_one','20d4574b-b5ee-8df6-5c05-184f1e29d493'),('93767ea8-6446-b574-c905-e40b60a38517','d36d9425-0d30-85f2-473c-1e0106d29c19','20d4574b-b5ee-8df6-5c05-184f1e29d493','user_one','SEMARANG, ISTORA SENAYAN','SENIN, 20 DESEMBER 2020','A','U','user_two','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9'),('b55b28a3-b281-14a8-6c68-c54bd453febd','d36d9425-0d30-85f2-473c-1e0106d29c19','20d4574b-b5ee-8df6-5c05-184f1e29d493','user_one','SEMARANG, ISTORA SENAYAN','SENIN, 20 DESEMBER 2020','A','U','user_two','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9'),('c827ccc6-a572-8c40-f7cd-9e15e1a87772','3ebb9d6a-f784-6f59-3a4c-761a4b7cdff0','20d4574b-b5ee-8df6-5c05-184f1e29d493','user_one','SEMARANG, ISTORA SENAYAN','SENIN, 20 DESEMBER 2020','I','W','user_two','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9'),('e0ce34f0-885f-e256-94a6-1d688183baef','d36d9425-0d30-85f2-473c-1e0106d29c19','20d4574b-b5ee-8df6-5c05-184f1e29d493','user_one','SEMARANG, ISTORA SENAYAN','SENIN, 20 DESEMBER 2020','I','W','user_two','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9'),('e377ccfa-49d9-cbc7-2ea5-15a1d685e818','ini ngaco','20d4574b-b5ee-8df6-5c05-184f1e29d493','user_one','SEMARANG, ISTORA SENAYAN','SENIN, 20 DESEMBER 2020','A','U','user_two','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9'),('f07070fb-525f-c14b-2413-c9744acffe83','029884f7-19ce-1e8a-8e1c-9a8236a02834','f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','user_two','awd','awdawd','A','U','user_one','20d4574b-b5ee-8df6-5c05-184f1e29d493');
/*!40000 ALTER TABLE `master_schedule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master_sport`
--

DROP TABLE IF EXISTS `master_sport`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `master_sport` (
  `sport_id` int NOT NULL AUTO_INCREMENT,
  `sport_name` varchar(50) DEFAULT NULL,
  `sport_player_count` int DEFAULT NULL,
  PRIMARY KEY (`sport_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master_sport`
--

LOCK TABLES `master_sport` WRITE;
/*!40000 ALTER TABLE `master_sport` DISABLE KEYS */;
INSERT INTO `master_sport` VALUES (1,'badminton single',2);
/*!40000 ALTER TABLE `master_sport` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master_user`
--

DROP TABLE IF EXISTS `master_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `master_user` (
  `user_id` varchar(50) NOT NULL,
  `user_photo_profile` longtext,
  `user_name` varchar(100) DEFAULT NULL,
  `user_full_name` varchar(100) DEFAULT NULL,
  `user_gender` varchar(1) DEFAULT NULL,
  `user_email` varchar(100) DEFAULT NULL,
  `user_password` varchar(100) DEFAULT NULL,
  `user_created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `user_facebook_account` varchar(200) NOT NULL,
  `user_google_account` varchar(200) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master_user`
--

LOCK TABLES `master_user` WRITE;
/*!40000 ALTER TABLE `master_user` DISABLE KEYS */;
INSERT INTO `master_user` VALUES ('20d4574b-b5ee-8df6-5c05-184f1e29d493','userPhoto-99247692554.jpg','user_one','user_one123','P','user_one@gmail.com','$2a$14$57eWkGE3kHjhFmJFgQyl0.fARv0h/X5hU8PT0pdZYooO5cBtbBrYu','2020-09-07 11:38:40','2020-09-07 11:38:40','',''),('296bf0dc-a0b6-9b35-9152-880296168dcf','userPhoto-17559938072.jpg','user_four','user_foure','L','userfoure@gmail.com','$2a$14$w5W5637z6HE/UVnhWj3p2Onm7ZQ9juM5OxZolWumyN2ktDLk4OlSG','2020-09-07 13:33:25','2020-09-07 13:33:25','adawd3','1231231211'),('9c267c5c-52bd-a55e-5d88-cc53a70128c8','userPhoto-17559938072.jpg','user_three','user_three','L','userthree@gmail.com','$2a$14$NOpyUCpf3HR70AGBSM/ky.55G90zSiRlZNH2/kucyXfoVvSl5KiNK','2020-09-07 12:26:22','2020-09-07 12:26:22','e12c12','23123123'),('ad208b3d-5159-3ad2-eccf-8da872bae25e','userPhoto-17559938072.jpg','admin','admin','L','admin@gmail.com','$2a$14$E0iYcL9r6VydjuP65j3RPOzLQqa2uXbdwoIf6wCtPoBuuOuHS4ACu','2020-09-09 20:52:40','2020-09-09 20:52:40','23412awd2123','222231'),('c031e52d-11f0-565d-237d-cf9a7497cd36','userPhoto-17559938072.jpg','buchori','buchori riyan','L','buchori@gmail.com','$2a$14$1SLCK0/5zFNSjb5aHdZNEeaUO1vE1dtLv1UlBpZavapckxOkqhI7i','2020-09-10 23:22:01','2020-09-10 23:22:01','1212d12','312'),('f0173ad5-89cb-ac87-c6d2-3ed9fdf438f9','userPhoto-43281744354.jpg','user_two','user_two123','L','usertwoe@gmail.com','$2a$14$/oGCkUtPMQxNtIbNPTIP4.ID4qHQcfCPka8DHujOsC4/Fa3ibSQ36','2020-09-07 11:57:53','2020-09-07 11:57:53','3782142135147313','107455961387022476700');
/*!40000 ALTER TABLE `master_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-09-20 18:42:28
