-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 08 Feb 2025 pada 03.52
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pusatdata`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `kecelakaan`
--

CREATE TABLE `kecelakaan` (
  `idKecelakaan` int(11) NOT NULL,
  `idKendaraan` int(11) NOT NULL,
  `penyebab` varchar(255) NOT NULL,
  `korban` varchar(255) NOT NULL,
  `tanggal` date NOT NULL,
  `waktu` datetime DEFAULT NULL,
  `lokasiKecelakaan` varchar(255) NOT NULL,
  `idLokasi` varchar(255) NOT NULL,
  `jenisJalur` enum('A','B','None Jalur') NOT NULL,
  `DeletedAt` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `kecelakaan`
--

INSERT INTO `kecelakaan` (`idKecelakaan`, `idKendaraan`, `penyebab`, `korban`, `tanggal`, `waktu`, `lokasiKecelakaan`, `idLokasi`, `jenisJalur`, `DeletedAt`) VALUES
(20, 2, 'Kurang Antisipasi', '2 Meninggal', '2025-02-15', '2025-02-14 22:38:44', '66-67', '8-PB', 'A', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `kendaraan`
--

CREATE TABLE `kendaraan` (
  `idKendaraan` int(11) NOT NULL,
  `namaKendaraan` varchar(255) NOT NULL,
  `warna` varchar(255) NOT NULL,
  `tipe` varchar(255) NOT NULL,
  `platNomor` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deletedAt` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `kendaraan`
--

INSERT INTO `kendaraan` (`idKendaraan`, `namaKendaraan`, `warna`, `tipe`, `platNomor`, `created_at`, `updated_at`, `deletedAt`) VALUES
(1, 'supra', 'putih', 'sport', '22233', '2025-01-12 06:18:45', '2025-01-24 05:42:49', '2025-01-24 12:42:49'),
(2, 'jeep', 'hitam', 'family', '232432', '2025-01-12 12:23:51', '2025-01-13 13:12:32', NULL),
(3, 'kijang', 'hitam', 'sport', '2332', '2025-01-12 16:39:24', '2025-01-13 13:12:39', NULL),
(4, 'bmw', 'putih', 'family', '32332', '2025-01-12 16:41:19', '2025-01-13 13:12:42', NULL),
(9, '1', '11', '111', '11', '2025-01-24 05:37:42', '2025-01-24 05:38:52', '2025-01-24 12:38:52'),
(10, 'minibus', '-', '-', '-', '2025-01-31 12:46:04', '2025-01-31 12:46:04', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `lokasi`
--

CREATE TABLE `lokasi` (
  `idLokasi` varchar(255) NOT NULL,
  `namaLokasi` varchar(255) NOT NULL,
  `mapLokasi` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `lokasi`
--

INSERT INTO `lokasi` (`idLokasi`, `namaLokasi`, `mapLokasi`, `created_at`, `updated_at`, `deleted_at`) VALUES
('1-PB', 'Kalihurip Itc. - Sadang Itc.', '-', '2025-01-13 11:02:37', '2025-01-13 11:02:37', NULL),
('10-PB', 'Kopo Itc - Moh.Toha Itc', '-', '2025-01-13 11:04:24', '2025-01-13 11:04:24', NULL),
('2-PB', 'Sadang Itc. - Jati Luhur Itc.', '-', '2025-01-13 11:02:56', '2025-01-13 11:02:56', NULL),
('3-PB', 'Sadang Itc. - Sadang', '-', '2025-01-13 11:03:07', '2025-01-13 11:03:07', NULL),
('4-PB', 'Jati Luhur Itc. - Padalarang Barat', '-', '2025-01-13 11:03:19', '2025-01-13 11:03:19', NULL),
('5-PB', 'RAM JT. LUHUR', '-', '2025-01-13 11:03:29', '2025-01-13 11:03:29', NULL),
('6-PB', 'Padalarang BRT- Padalarang', '-', '2025-01-13 11:03:39', '2025-01-13 11:03:39', NULL),
('7-PB', 'Padalarang - Pasteur Itc', '-', '2025-01-13 11:03:50', '2025-01-13 11:03:50', NULL),
('8-PB', 'Pasteur Itc - Pasir Koja Itc', '-', '2025-01-13 11:04:03', '2025-01-13 11:04:03', NULL),
('9-PB', 'Pasir Koja Itc - Kopo Itc', '-', '2025-01-13 11:04:12', '2025-01-13 11:04:12', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `pengaduan`
--

CREATE TABLE `pengaduan` (
  `id_pengaduan` int(11) NOT NULL,
  `tanggal_waktu` datetime NOT NULL,
  `lokasi_kecelakaan` varchar(255) NOT NULL,
  `id_kendaraan` varchar(11) NOT NULL,
  `jumlah_kendaraan` int(11) NOT NULL,
  `id_lokasi` varchar(255) NOT NULL,
  `jenis_jalur` enum('A','B') NOT NULL,
  `cuaca` varchar(255) NOT NULL,
  `jalur_tertutup_total` enum('ya','tidak') NOT NULL,
  `status_pengaduan` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `pengaduan`
--

INSERT INTO `pengaduan` (`id_pengaduan`, `tanggal_waktu`, `lokasi_kecelakaan`, `id_kendaraan`, `jumlah_kendaraan`, `id_lokasi`, `jenis_jalur`, `cuaca`, `jalur_tertutup_total`, `status_pengaduan`) VALUES
(1, '2025-02-06 09:30:00', '68-69', '2', 12, '1-PB', 'A', '2', 'ya', 'valid'),
(2, '2025-02-06 09:30:00', '2', '2', 2, '1-PB', 'A', '2', 'ya', 'belum_ditanggapi'),
(3, '2025-02-08 09:52:00', '69-70', '2', 2, '1-PB', 'A', '2', 'ya', 'belum_ditanggapi'),
(4, '2025-02-08 09:52:00', '69-70', '23', 2, '1-PB', 'A', '2', 'ya', 'belum_ditanggapi'),
(5, '2025-02-08 07:55:00', '68-69', '232', 23, '1-PB', 'A', '23', 'ya', 'valid'),
(6, '2025-02-08 10:21:00', '68-69', '123', 123, '10-PB', 'A', '123', 'ya', 'valid'),
(7, '2025-02-08 09:34:00', '66-67', '123', 123, '8-PB', 'A', '42', 'tidak', 'belum_ditanggapi'),
(8, '2025-02-08 10:35:00', '66-67', '23', 23, '7-PB', 'A', '23', 'ya', 'valid'),
(9, '2025-02-08 09:46:00', '83-84', '23', 23, '5-PB', 'A', '2', 'ya', 'belum_ditanggapi');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(50) DEFAULT 'user',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password`, `role`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'sanjaya', 'sanjaya@gmail.com', 'sanjaya', 'admin', '2025-01-07 08:47:08', '2025-01-24 06:59:22', NULL),
(8, 'user', 'user@gmail.com', 'users', 'user', '2025-01-10 07:01:48', '2025-01-24 06:59:26', NULL);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `kecelakaan`
--
ALTER TABLE `kecelakaan`
  ADD PRIMARY KEY (`idKecelakaan`);

--
-- Indeks untuk tabel `kendaraan`
--
ALTER TABLE `kendaraan`
  ADD PRIMARY KEY (`idKendaraan`);

--
-- Indeks untuk tabel `lokasi`
--
ALTER TABLE `lokasi`
  ADD PRIMARY KEY (`idLokasi`),
  ADD UNIQUE KEY `nama_lokasi` (`namaLokasi`);

--
-- Indeks untuk tabel `pengaduan`
--
ALTER TABLE `pengaduan`
  ADD PRIMARY KEY (`id_pengaduan`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `kecelakaan`
--
ALTER TABLE `kecelakaan`
  MODIFY `idKecelakaan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT untuk tabel `kendaraan`
--
ALTER TABLE `kendaraan`
  MODIFY `idKendaraan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT untuk tabel `pengaduan`
--
ALTER TABLE `pengaduan`
  MODIFY `id_pengaduan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
