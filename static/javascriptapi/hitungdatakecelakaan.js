function fetchTotalKecelakaan() {
    fetch('/api/kecelakaan/total')
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to fetch total kecelakaan');
            }
            return response.json();
        })
        .then(data => {
            // Periksa apakah data memiliki properti 'total'
            if (data && typeof data.total === 'number') {
                // Update elemen H5 dengan jumlah kecelakaan
                document.querySelector('h5.font-weight-bolder').textContent = data.total + " Kejadian";
            } else {
                console.error('Invalid response data:', data);
            }
        })
        .catch(error => {
            console.error('Error fetching kecelakaan data:', error);
        });
}

// Panggil fungsi pada saat halaman dimuat
fetchTotalKecelakaan();
function decodeBase64(base64String) {
    return atob(base64String);
}

function fetchTotalKecelakaanMobilTahunIni() {
    fetch('/api/kecelakaan/mobil')
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to fetch total kecelakaan');
            }
            return response.json();
        })
        .then(data => {
            // Periksa apakah data memiliki properti 'tipeKendaraan' dan 'total'
            if (data && data.tipeKendaraan && typeof data.total === 'number') {
                // Decode tipe kendaraan jika dalam base64
                const decodedTipeKendaraan = decodeBase64(data.tipeKendaraan);

                // Update elemen H5 dengan tipe kendaraan yang sudah di-decode dan jumlah kecelakaan
                document.querySelector('#textmobil').textContent = `Tipe kendaraan: ${decodedTipeKendaraan} - Total kecelakaan: ${data.total}`;
            } else {
                console.error('Invalid response data:', data);
            }
        })
        .catch(error => {
            console.error('Error fetching kecelakaan data:', error);
        });
}

// Panggil fungsi pada saat halaman dimuat
fetchTotalKecelakaanMobilTahunIni();



function fetchTotalKecelakaanPerLokasi() { 
    fetch('/api/kecelakaan/per-lokasi') // Ganti endpoint sesuai dengan API Anda
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to fetch total kecelakaan per lokasi');
            }
            return response.json();
        })
        .then(data => {
            // Pastikan data adalah array
            if (Array.isArray(data)) {
                const container = document.querySelector('#lokasiKecelakaanList'); // Container untuk menampilkan data

                // Kosongkan container sebelum menambahkan data baru
                container.innerHTML = '';

                // Temukan item dengan total kecelakaan tertinggi
                const maxItem = data.reduce((max, item) => {
                    return (item.totalKecelakaan > max.totalKecelakaan) ? item : max;
                }, data[0]); // Asumsikan data tidak kosong, jika kosong perlu pengecekan lebih lanjut

                // Tampilkan lokasi dengan kecelakaan terbanyak
                if (maxItem.namaLokasi && typeof maxItem.totalKecelakaan === 'number') {
                    const lokasiItem = document.createElement('div');
                    lokasiItem.textContent = `${maxItem.namaLokasi}: ${maxItem.totalKecelakaan} kecelakaan`;
                    container.appendChild(lokasiItem);
                } else {
                    console.error('Invalid item data:', maxItem);
                }
            } else {
                console.error('Invalid response data:', data);
            }
        })
        .catch(error => {
            console.error('Error fetching kecelakaan data:', error);
        });
}



// Panggil fungsi pada saat halaman dimuat
fetchTotalKecelakaanPerLokasi();



function fetchPenyebabTertinggi() { 
    fetch('/api/kecelakaan/penyebab-tertinggi') 
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to fetch penyebab data');
            }
            return response.json();
        })
        .then(data => {
            if (data && data.penyebab && data.count) { 
                // Update elemen H5 dengan penyebab kecelakaan terbanyak
                document.querySelector('#textPenyebab').textContent = 
                    `${data.penyebab} , ${data.count} kejadian`;
            } else {
                console.error('Invalid response data:', data);
            }
        })
        .catch(error => {
            console.error('Error fetching penyebab data:', error);
        });
}

// Panggil fungsi pada saat halaman dimuat
fetchPenyebabTertinggi();
