-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1:3306
-- Tiempo de generación: 23-10-2023 a las 14:39:22
-- Versión del servidor: 5.7.31
-- Versión de PHP: 7.4.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `tareas_plus`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `categorias`
--

DROP TABLE IF EXISTS `categorias`;
CREATE TABLE IF NOT EXISTS `categorias` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Nombre` varchar(255) COLLATE latin1_spanish_ci NOT NULL,
  `Color` varchar(20) COLLATE latin1_spanish_ci DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

--
-- Volcado de datos para la tabla `categorias`
--

INSERT INTO `categorias` (`ID`, `Nombre`, `Color`) VALUES
(1, 'PRUEBAA', 'PRUEBAA'),
(2, 'Importante', 'Azul'),
(4, 'Importante_rojo_4', 'Rojo_4');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `historialtareas`
--

DROP TABLE IF EXISTS `historialtareas`;
CREATE TABLE IF NOT EXISTS `historialtareas` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `AccionRealizada` text COLLATE latin1_spanish_ci NOT NULL,
  `FechaHoraAccion` datetime NOT NULL,
  `ID_Usuario` int(11) DEFAULT NULL,
  `ID_Tarea` int(11) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_Usuario` (`ID_Usuario`),
  KEY `ID_Tarea` (`ID_Tarea`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `notificaciones`
--

DROP TABLE IF EXISTS `notificaciones`;
CREATE TABLE IF NOT EXISTS `notificaciones` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Descripcion` text COLLATE latin1_spanish_ci NOT NULL,
  `FechaHoraNotificacion` datetime NOT NULL,
  `ID_Tarea` int(11) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_Tarea` (`ID_Tarea`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `prioridades`
--

DROP TABLE IF EXISTS `prioridades`;
CREATE TABLE IF NOT EXISTS `prioridades` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Nombre` varchar(50) COLLATE latin1_spanish_ci NOT NULL,
  `Descripcion` text COLLATE latin1_spanish_ci,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `tareas`
--

DROP TABLE IF EXISTS `tareas`;
CREATE TABLE IF NOT EXISTS `tareas` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Titulo` varchar(255) COLLATE latin1_spanish_ci NOT NULL,
  `Descripcion` text COLLATE latin1_spanish_ci,
  `fecha_vencimiento` datetime DEFAULT NULL,
  `Estado` enum('Pendiente','Completada') COLLATE latin1_spanish_ci NOT NULL,
  `ID_Categoria` int(11) DEFAULT NULL,
  `ID_Prioridad` int(11) DEFAULT NULL,
  `ID_Usuario` int(11) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `ID_Categoria` (`ID_Categoria`),
  KEY `ID_Prioridad` (`ID_Prioridad`),
  KEY `ID_Usuario` (`ID_Usuario`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `usuarios`
--

DROP TABLE IF EXISTS `usuarios`;
CREATE TABLE IF NOT EXISTS `usuarios` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Nombre` varchar(255) COLLATE latin1_spanish_ci NOT NULL,
  `CorreoElectronico` varchar(255) COLLATE latin1_spanish_ci NOT NULL,
  `Contrasena` varchar(255) COLLATE latin1_spanish_ci NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COLLATE=latin1_spanish_ci;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
