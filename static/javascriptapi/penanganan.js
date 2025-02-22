let currentPage = 1;
const limit = 7;

function fetchAndLoadPengaduan() {
    // Gunakan variabel currentPage dan limit dalam URL
    const url = `/api/Pengaduan/?page=${currentPage}&limit=${limit}`;
    fetch(url)
      .then(response => {
         if (!response.ok) {
             throw new Error('Failed to fetch data');
         }
         return response.text(); // Ambil respons sebagai text
      })
      .then(rawText => {
         console.log('Raw response:', rawText);
         // Jika terdapat dua objek JSON yang tergabung, pisahkan dan gabungkan (merge)
         const parts = rawText.split('}{');
         let data;
         if (parts.length >= 2) {
             // Tambahkan kembali kurung kurawal yang hilang
             const firstPart = parts[0] + '}';
             const secondPart = '{' + parts[1];
             try {
                 const json1 = JSON.parse(firstPart);
                 const json2 = JSON.parse(secondPart);
                 // Gabungkan kedua objek. Jika ada properti yang sama, gunakan nilai dari json2.
                 data = Object.assign({}, json1, json2);
             } catch (err) {
                 console.error('Error parsing merged JSON:', err);
                 throw err;
             }
         } else {
             try {
                 data = JSON.parse(rawText);
             } catch (err) {
                 console.error('Error parsing JSON:', err);
                 throw err;
             }
         }
         return data;
      })
      .then(data => {
         console.log('Response Data:', data); // Log untuk debugging

         if (data && data.data) {
             let pengaduanData = data.data;

             // Log data sebelum sorting
             console.log('Before Sorting:', pengaduanData.map(item => item.id_pengaduan));

             // Sorting berdasarkan id_pengaduan
             pengaduanData.sort((a, b) => a.id_pengaduan - b.id_pengaduan);

             // Log data setelah sorting
             console.log('After Sorting:', pengaduanData.map(item => item.id_pengaduan));

             const tableData = pengaduanData.map((pengaduan, index) => ({
                 'ID Pengaduan': pengaduan.id_pengaduan,
                 'Tanggal Waktu': pengaduan.tanggal_waktu,
                 'Lokasi Kecelakaan': pengaduan.lokasi_kecelakaan,
                 'ID Kendaraan': pengaduan.id_kendaraan,
                 'Jumlah Kendaraan': pengaduan.jumlah_kendaraan,
                 'Lokasi Peruas': pengaduan.id_lokasi,
                 'Jenis Jalur': pengaduan.jenis_jalur,
                 'Cuaca': pengaduan.cuaca,
                 'Jalur Tertutup Total': pengaduan.jalur_tertutup_total,
                 'Status Pengaduan': pengaduan.status_pengaduan,
                 'Actions': index + 1
             }));

             // Memuat data ke UI
             loadTableD24(
                 ['ID Pengaduan', 'Tanggal Waktu', 'Lokasi Kecelakaan', 'ID Kendaraan', 'Jumlah Kendaraan', 'Lokasi Peruas', 'Jenis Jalur', 'Cuaca', 'Jalur Tertutup Total', 'Status Pengaduan', 'Actions'],
                 tableData
             );

             // Menangani pagination
             handlePaginationPengaduan(data.totalPages);
         } else {
             console.error('Invalid response data:', data);
             alert('Gagal menampilkan data');
         }
      })
      .catch(error => {
         console.error('Error fetching Pengaduan data:', error);
         alert('Gagal memuat data Pengaduan');
      });
}

function handlePaginationPengaduan(totalPages) {
    const paginationContainer = document.getElementById('pagination-container');
    if (!paginationContainer) {
      console.error('Pagination container not found');
      return;
    }
  
    let paginationHTML = '';
    for (let i = 1; i <= totalPages; i++) {
      paginationHTML += `<button class="btn btn-sm btn-secondary" style="margin-right:10px;" onclick="goToPagePengaduan(${i})">${i}</button>`;
    }
  
    paginationContainer.innerHTML = paginationHTML;
}
  
function goToPagePengaduan(pageNumber) {
    currentPage = pageNumber;
    fetchAndLoadPengaduan();
}


  
// Fungsi untuk memuat data tabel ke dalam UI
function loadTableD24(headers, data) {
    console.log('Table Headers:', headers);
    console.log('Table Data:', data);

    const tableContainer = document.getElementById('table-container');
    if (!tableContainer) {
        console.error('Table container not found');
        return;
    }

    let tableHTML = '<thead><tr>';
    headers.forEach(header => {
        tableHTML += `<th>${header}</th>`;
    });
    tableHTML += '</tr></thead><tbody>';

    data.forEach((row) => {
        tableHTML += '<tr>';
        headers.forEach(header => {
            // Jika header adalah 'Actions', tambahkan tombol aksi
            if (header === 'Actions') {
                tableHTML += `
                    <td>
                        <button class="btn btn-sm btn-warning" onclick="beritanggapantidakvalid('${row['ID Pengaduan']}')">tidak valid</button>
                        <button class="btn btn-sm btn-success" onclick="beritanggapanvalid('${row['ID Pengaduan']}')">valid</button>
                        <button class="btn btn-sm btn-danger" onclick="deletepenanganRow('${row['ID Pengaduan']}')">Delete</button>
                    </td>
                `;
            } else {
                tableHTML += `<td>${row[header] || ''}</td>`;
            }
        });
        tableHTML += '</tr>';
    });

    tableHTML += '</tbody>';
    tableContainer.innerHTML = tableHTML;
}


function beritanggapantidakvalid(idPengaduan) {
    if (confirm('Apakah anda ingin memberi tanggapan tidak valid ?')) {
        fetch(`/api/Pengaduan/${idPengaduan}/tidakvalid`, {
                method: 'POST'
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`gagal mengubah ID ${idPengaduan}`);
                }
                return response.json();
            })
            .then(data => {
                console.log('Response Data:', data); // Log for debugging
                if (data.message && data.message === "Pengaduan berhasil diperbarui") {
                    alert('Pengaduan berhasil diperbarui');
                    fetchAndLoadPengaduan(); // Reload table data
                } else {
                    alert('Pengaduan gagal diperbarui');
                }
            })
            .catch(error => {
                console.error('Error  Kecelakaan:', error);
                alert('An error occurred while  the record.');
            });
    }
}

function beritanggapanvalid(idPengaduan) {
    if (confirm('Apakah anda ingin memberi tanggapan valid ?')) {
        fetch(`/api/Pengaduan/${idPengaduan}/valid`, {
                method: 'POST'
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`gagal mengubah  ID ${idPengaduan}`);
                }
                return response.json();
            })
            .then(data => {
                console.log('Response Data:', data); // Log for debugging
                if (data.message && data.message === "Pengaduan berhasil diperbarui") {
                    alert('Pengaduan berhasil diperbarui ');
                    fetchAndLoadPengaduan(); // Reload table data
                } else {
                    alert('Pengaduan gagal diperbarui');
                }
            })

            .catch(error => {
                console.error('Error  pengaduan:', error);
                alert('An error occurred while  the record.');
            });
    }
}



function deletepenanganRow(idpenangan) {
    if (confirm('Are you sure you want to delete this penangan?')) {
        fetch(`/api/penangan/${idpenangan}`, {
                method: 'DELETE'
            })
            .then(response => response.json())
            .then(data => {
                if (data.message === "Pengaduan soft-deleted successfully") {
                    alert('penangan deleted successfully');
                    fetchAndLoadPengaduan(); // Reload table data
                } else {
                    alert('Failed to delete penangan');
                }
            })
            .catch(error => console.error('Error deleting penangan:', error));
    }
}
// Now call the function to load data
fetchAndLoadPengaduan();