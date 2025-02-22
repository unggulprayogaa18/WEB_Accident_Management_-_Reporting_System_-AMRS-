///////////////////////////////////////////////////////////////////////////// LOKASI ///////////////////////////////////////////////////////////////////////////////////////
let currentPage = 1;
const limit = 7;

// Load default form and table on page load
document.addEventListener('DOMContentLoaded', function () {
  document.getElementById('form-kecelakaan-tab').click();
});

// Fungsi untuk mendapatkan ID Lokasi dari API
function getLokasiID() {
  axios.get('/api/generate-lokasi-id')
    .then(response => {
      console.log('Response from API:', response);
      // Menampilkan ID Lokasi pada input field dengan id 'idLokasi'
      document.getElementById('idLokasi').value = response.data.idLokasi;
    })
    .catch(error => {
      console.error('Error occurred while fetching Lokasi ID:', error);
    });
}


// Memanggil API untuk mendapatkan Lokasi ID saat halaman dimuat
document.addEventListener('DOMContentLoaded', function () {
  getLokasiID();
});


// Event listener untuk menampilkan form lokasi ketika tab diklik
document.getElementById('form-lokasi-tab').addEventListener('click', function () {
  document.getElementById('form-container2').style.display = 'none';
  document.getElementById('form-container3').style.display = 'none';

  // Show form-container2 and insert dynamic content into it
  document.getElementById('form-container1').style.display = 'block';
  document.getElementById('form-container1').innerHTML = `
     <form id="lokasiForm">
       <div class="mb-3">
         <label for="idLokasi" class="form-label">ID Lokasi</label>
         <input type="text" class="form-control" id="idLokasi" readonly>
       </div>
   <div class="mb-3">
     <label for="namaLokasi" class="form-label">Nama Lokasi</label>
     <input type="text" class="form-control" id="namaLokasi" placeholder="Enter Name">
   </div>
   <div class="mb-3">
     <label for="mapLokasi" class="form-label">Map Lokasi</label>
     <input type="text" class="form-control" id="mapLokasi" placeholder="Enter Map Location">
   </div>
   <button type="submit" class="btn btn-primary">Submit</button>
 </form>

   `;

  // Setelah form dimuat, panggil kembali fungsi untuk mendapatkan ID Lokasi
  getLokasiID();
  // Panggil API untuk memuat data Lokasi ke tabel
  fetchAndLoadLokasi();

  document.getElementById('lokasiForm').addEventListener('submit', async function (e) {
    e.preventDefault(); // Prevent the default form submission behavior

    // Ambil nilai dari input form
    const idLokasi = document.getElementById('idLokasi').value;
    const namaLokasi = document.getElementById('namaLokasi').value;
    const mapLokasi = document.getElementById('mapLokasi').value;

    // Debug log untuk memastikan nilai yang diambil dari form
    console.log('Form values:', {
      idLokasi,
      namaLokasi,
      mapLokasi
    });

    // Validasi input form (optional)
    if (!idLokasi || !namaLokasi || !mapLokasi) {
      console.log('Form validation failed: All fields must be filled');
      alert('Semua kolom harus diisi');
      return;
    }

    // Kirim data ke endpoint API CreateLokasi
    try {
      console.log('Sending request to /api/lokasi/');

      const response = await fetch('/api/lokasi/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          idLokasi: idLokasi,
          namaLokasi: namaLokasi,
          mapLokasi: mapLokasi,
        }),
      });

      // Debug log untuk status response
      console.log('Response status:', response.status);

      if (!response.ok) {
        throw new Error('Failed to create Lokasi');
      }

      const result = await response.json();

      // Debug log untuk hasil respon dari server
      console.log('Server response:', result);

      alert('Lokasi created successfully: ' + JSON.stringify(result));
      // Setelah form dimuat, panggil kembali fungsi untuk mendapatkan ID Lokasi
      getLokasiID();

      // Refresh halaman setelah berhasil mendapatkan ID
      window.location.reload();
    } catch (error) {
      // Log error yang terjadi
      console.log('Error occurred:', error);
      alert('Error: ' + error.message);
    }
  });




  // Setelah form dimuat, panggil kembali fungsi untuk mendapatkan ID Lokasi
  getLokasiID();


});



function fetchAndLoadLokasi() {
  // Panggil API untuk mendapatkan data Lokasi dengan pagination
  axios.get(`/api/lokasi/?page=${currentPage}&limit=${limit}`)
    .then(response => {
      console.log('Response from GetAllLokasi:', response);

      if (response.data && response.data.data) {
        let lokasiData = response.data.data;

        // Mengurutkan berdasarkan angka di bagian awal 'idLokasi'
        lokasiData.sort((a, b) => {
          const numA = parseInt(a.idLokasi.split('-')[0], 10);
          const numB = parseInt(b.idLokasi.split('-')[0], 10);
          return numA - numB; // Urutkan berdasarkan angka
        });

        const tableData = lokasiData.map((lokasi, index) => ({
          'ID Lokasi': lokasi.idLokasi,
          'Nama Lokasi': lokasi.namaLokasi,
          'Map Lokasi': lokasi.mapLokasi,
          'Actions': index + 1 // Starting index from 1 instead of 0
        }));

        loadTableD(
          ['ID Lokasi', 'Nama Lokasi', 'Map Lokasi', 'Actions'], // Add 'Actions' header
          tableData
        );

        // Handle pagination
        handlePaginationlokasi(response.data.totalPages);
      } else {
        console.error('Invalid response data:', response.data);
      }
    })
    .catch(error => {
      console.error('Error fetching Lokasi data:', error);
      alert('Gagal memuat data Lokasi');
    });
}

function handlePaginationlokasi(totalPages) {
  const paginationContainer = document.getElementById('pagination-container');
  if (!paginationContainer) {
    console.error('Pagination container not found');
    return;
  }

  let paginationHTML = '';
  for (let i = 1; i <= totalPages; i++) {
    paginationHTML += `<button class="btn btn-sm btn-secondary" style="margin-right:10px;" onclick="goToPageLokasi(${i})">${i}</button>`;
  }

  paginationContainer.innerHTML = paginationHTML;
}

function goToPageLokasi(pageNumber) {
  currentPage = pageNumber;
  fetchAndLoadLokasi();
}

// Call fetchAndLoadLokasi initially

function loadTableD(headers, data) {
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
      // If the header is 'Actions', add the action buttons
      if (header === 'Actions') {
        tableHTML += `
                   <td>
                       <button class="btn btn-sm btn-primary" onclick="editRowLokasi('${row['ID Lokasi']}')">Edit</button>
                       <button class="btn btn-sm btn-danger" onclick="deleteRowLokasi('${row['ID Lokasi']}')">Delete</button>
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

function editRowLokasi(lokasiId) {
  // Memanggil API untuk mendapatkan data lokasi berdasarkan ID
  fetch(`/api/lokasi/${lokasiId}`)
    .then(response => response.json())
    .then(data => {
      if (data.data) {
        // Isi form dengan data lokasi
        document.getElementById('id_lokasi').value = data.data.idLokasi; // Using 'idLokasi' from response
        document.getElementById('nama_lokasi').value = data.data.namaLokasi;
        document.getElementById('map_lokasi').value = data.data.mapLokasi;

        // Tampilkan overlay
        toggleOverlay(1);
      } else {
        console.error('Lokasi not found');
      }
    })
    .catch(error => {
      console.error('Error fetching lokasi:', error);
    });
}

function updateLokasi() {
  // Ambil data dari form
  const idLokasi = document.getElementById('id_lokasi').value;
  const namaLokasi = document.getElementById('nama_lokasi').value;
  const mapLokasi = document.getElementById('map_lokasi').value;

  // Kirimkan data ke API untuk update
  fetch(`/api/lokasi/${idLokasi}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        // Ubah nama properti untuk mencocokkan nama yang diharapkan oleh backend
        namaLokasi: namaLokasi,
        mapLokasi: mapLokasi,
      }),
    })
    .then(response => response.json())
    .then(data => {
      if (data.message === "Lokasi updated successfully") {
        alert('Data Lokasi berhasil diperbarui');
        toggleOverlay(1); // Tutup overlay setelah berhasil update
        // Refresh atau update tampilan data di halaman jika diperlukan
        fetchAndLoadLokasi();
      } else {
        alert('Gagal memperbarui data');
      }
    })
    .catch(error => {
      console.error('Error updating lokasi:', error);
    });
}

function deleteRowLokasi(lokasiId) {
  if (confirm('Are you sure you want to delete this Lokasi?')) {
    // Send DELETE request to API
    fetch(`/api/lokasi/${lokasiId}`, {
        method: 'DELETE',
      })
      .then(response => response.json())
      .then(data => {
        if (data.message === "Lokasi deleted successfully") {
          alert('Data Lokasi berhasil dihapus');
          // Refresh or update table data
          fetchAndLoadLokasi();
        } else {
          alert('Gagal menghapus data Lokasi');
        }
      })
      .catch(error => {
        console.error('Error deleting lokasi:', error);
      });
  }
}

/////////////////////////////////////////////////////////////////////////// Kecelakaan ///////////////////////////////////////////////////////////////////////////////////////

// Event listener for tab click
document.getElementById('form-kecelakaan-tab').addEventListener('click', function () {
  document.getElementById('form-container1').style.display = 'none';
  document.getElementById('form-container2').style.display = 'none';

  // Show form-container3 and insert dynamic content into it
  document.getElementById('form-container3').style.display = 'block';

  // Fetch data from backend using Axios
  axios.get('/ambildatakendaranlokasi/getformdata')
    .then(function (response) {
      // Log the response to check if the data is correct
      console.log("Response from API:", response);

      // Extract the data from the response
      const kendaraanData = response.data.kendaraan;
      const lokasiData = response.data.lokasi;

      // Log the extracted data
      console.log("Kendaraan Data:", kendaraanData);
      console.log("Lokasi Data:", lokasiData);

      // Build the HTML for the form with comboboxes using template literals
      let formHTML = `
      <form id="kecelakaanForm">

       <div class="mb-3">
            <label for="penyebab" class="form-label">Penyebab Kecelakaan</label>
            <input type="text" class="form-control" id="penyebab" placeholder="Enter Cause">
          </div>
          <div class="mb-3">
            <label for="korban" class="form-label">Jumlah Korban</label>
            <input type="text" class="form-control" id="korban" placeholder="Enter Speed">
          </div>
          <div class="mb-3">
            <label for="tanggal" class="form-label">Tanggal Kejadian</label>
            <input type="date" class="form-control" id="tanggal">
          </div>
          <div class="mb-3">
            <label for="waktu" class="form-label">Waktu Kejadian</label>
            <input type="time" class="form-control" id="waktu">
          </div>

          <div class="mb-3">
            <label for="lokasikecelekaan" class="form-label">Lokasi Kecelakaan (km)</label>
             <select id="lokasikecelekaan" class="form-select">
              <option value="66-67">66 - 67 KM</option>
               <option value="67-68">67 - 68 KM</option>
               <option value="68-69">68 - 69 KM</option>
               <option value="69-70">69 - 70 KM</option>
               <option value="70-71">70 - 71 KM</option>
               <option value="71-72">71 - 72 KM</option>
               <option value="72-73">72 - 73 KM</option>
               <option value="73-74">73 - 74 KM</option>
               <option value="74-75">74 - 75 KM</option>
               <option value="75-76">75 - 76 KM</option>
               <option value="76-77">76 - 77 KM</option>
               <option value="77-78">77 - 78 KM</option>
               <option value="78-79">78 - 79 KM</option>
               <option value="79-80">79 - 80 KM</option>
               <option value="80-81">80 - 81 KM</option>
               <option value="81-82">81 - 82 KM</option>
               <option value="82-83">82 - 83 KM</option>
               <option value="83-84">83 - 84 KM</option>
               <option value="84-85">84 - 85 KM</option>
               <option value="85-86">85 - 86 KM</option>
               <option value="86-87">86 - 87 KM</option>
               <option value="87-88">87 - 88 KM</option>
               <option value="88-89">88 - 89 KM</option>
               <option value="89-90">89 - 90 KM</option>
               <option value="90-91">90 - 91 KM</option>
               <option value="91-92">91 - 92 KM</option>
               <option value="92-93">92 - 93 KM</option>
               <option value="93-94">93 - 94 KM</option>
               <option value="94-95">94 - 95 KM</option>
               <option value="95-96">95 - 96 KM</option>
               <option value="96-97">96 - 97 KM</option>
               <option value="97-98">97 - 98 KM</option>
               <option value="98-99">98 - 99 KM</option>
               <option value="99-100">99 - 100 KM</option>
               <option value="100-101">100 - 101 KM</option>
               <option value="101-102">101 - 102 KM</option>
               <option value="102-103">102 - 103 KM</option>
               <option value="103-104">103 - 104 KM</option>
               <option value="104-105">104 - 105 KM</option>
               <option value="105-106">105 - 106 KM</option>
               <option value="106-107">106 - 107 KM</option>
               <option value="107-108">107 - 108 KM</option>
               <option value="108-109">108 - 109 KM</option>
               <option value="109-110">109 - 110 KM</option>
               <option value="110-111">110 - 111 KM</option>
               <option value="111-112">111 - 112 KM</option>
               <option value="112-113">112 - 113 KM</option>
               <option value="113-114">113 - 114 KM</option>
               <option value="114-115">114 - 115 KM</option>
               <option value="115-116">115 - 116 KM</option>
               <option value="116-117">116 - 117 KM</option>
               <option value="117-118">117 - 118 KM</option>
               <option value="118-119">118 - 119 KM</option>
               <option value="119-120">119 - 120 KM</option>
               <option value="120-121">120 - 121 KM</option>
               <option value="121-122">121 - 122 KM</option>
               <option value="122-123">122 - 123 KM</option>
               <option value="123-124">123 - 124 KM</option>
               <option value="124-125">124 - 125 KM</option>
               <option value="125-126">125 - 126 KM</option>
               <option value="126-127">126 - 127 KM</option>
               <option value="127-128">127 - 128 KM</option>
               <option value="128-129">128 - 129 KM</option>
               <option value="129-130">129 - 130 KM</option>
               <option value="130-131">130 - 131 KM</option>
               <option value="131-132">131 - 132 KM</option>
               <option value="132-133">132 - 133 KM</option>
               <option value="133-134">133 - 134 KM</option>
               <option value="134-135">134 - 135 KM</option>
               <option value="135-136">135 - 136 KM</option>
               <option value="136-137">136 - 137 KM</option>
               <option value="137-138">137 - 138 KM</option>
               <option value="138-139">138 - 139 KM</option>
               <option value="139-140">139 - 140 KM</option>
               <option value="140-141">140 - 141 KM</option>
               <option value="141-142">141 - 142 KM</option>
               <option value="142-143">142 - 143 KM</option>
               <option value="143-144">143 - 144 KM</option>
               <option value="144-145">144 - 145 KM</option>
               <option value="145-146">145 - 146 KM</option>
               <option value="146-147">146 - 147 KM</option>
               <option value="147-148">147 - 148 KM</option>
               <option value="148-149">148 - 149 KM</option>
               <option value="149-150">149 - 150 KM</option>
               <option value="150-151">150 - 151 KM</option>
               <option value="151-152">151 - 152 KM</option>
               <option value="152-153">152 - 153 KM</option>
               <option value="153-154">153 - 154 KM</option>
               <option value="154-155">154 - 155 KM</option>
               <option value="155-156">155 - 156 KM</option>
               <option value="156-157">156 - 157 KM</option>
               <option value="00-01-BR">00-01 BR KM</option>
               <option value="01-02-PS">01-02 PS KM</option>
               <option value="02-03-PS">02-03 PS KM</option>
               <option value="03-04-PS">03-04 PS KM</option>
               <option value="04-05-PS">04-05 PS KM</option>
               <option value="05-06-PS">05-06 PS KM</option>
               <option value="00-01-KJ">00-01 KJ KM</option>
               <option value="01-02-KJ">01-02 KJ KM</option>
               <option value="02-03-KJ">02-03 KJ KM</option>
          </select>
          </div>

          <div class="mb-3">
            <input type="checkbox" id="otherLocationCheckbox"> Pilih lokasi kecelakaan lainnya
          </div>

          <!-- Bagian untuk lokasi lain yang akan muncul jika checkbox dipilih -->
          <div class="mb-3" id="otherLocationContainer" style="display: none;">
            <label for="otherLocation" class="form-label">Lokasi Lain</label>
            <input type="text" class="form-control" id="otherLocation" placeholder="Masukkan Lokasi Lain">
          </div>

        <div class="mb-3">
          <label for="combokendaran">Select Kendaraan:</label>
          <select id="combokendaran" class="form-control" name="idKendaraan">
            <option value="">Select Kendaraan</option>
            ${kendaraanData.map(kendaraan => `<option value="${kendaraan.idKendaraan}">${kendaraan.idKendaraan} - ${kendaraan.namaKendaraan}</option>`).join('')}
          </select>
        </div>

        <div class="mb-3">
          <label for="lokasiPeruas">Select Lokasi:</label>
          <select id="lokasiPeruas" class="form-control" name="idLokasi">
            <option value="">Select Lokasi per ruas</option>
            ${lokasiData.map(lokasi => `<option value="${lokasi.idLokasi}">${lokasi.idLokasi} - ${lokasi.namaLokasi}</option>`).join('')}
          </select>
        </div>


        <div class="mb-3">
        <label for="jenisJalur" class="form-label">Jenis Jalur</label>
        <div>

        <input type="radio" id="jenisJalurA" name="jenisJalur" value="A">
        <label for="jenisJalurA">A (Bandung)</label>
        </div>
        
        <div>
        <input type="radio" id="jenisJalurB" name="jenisJalur" value="B">
        <label for="jenisJalurB">B (jakarta)</label>
        </div>
        
        <div>
        <input type="radio" id="jenisJalurNone" name="jenisJalur" value="Bahu Jalur">
        <label for="jenisJalurNone">Bahu Jalur</label>
        </div>
        
        </div>

        <button type="submit" class="btn btn-primary">Submit</button>
      </form>
    `;

      fetchAndLoadKecelakaan();
      // Insert the generated HTML into form-container3
      document.getElementById('form-container3').innerHTML = formHTML;

      // Add event listener for checkbox to toggle the location input field
      document.getElementById('otherLocationCheckbox').addEventListener('change', function () {
        var otherLocationContainer = document.getElementById('otherLocationContainer');
        var lokasiSelect = document.getElementById('lokasikecelekaan');

        if (this.checked) {
          // Hide the location select and show the text input field
          lokasiSelect.style.display = 'none';
          otherLocationContainer.style.display = 'block';
        } else {
          // Show the location select and hide the text input field
          lokasiSelect.style.display = 'block';
          otherLocationContainer.style.display = 'none';
        }
      });


      document.getElementById('kecelakaanForm').addEventListener('submit', function (event) {
        event.preventDefault();

        // Validasi dan persiapan data
        const penyebab = document.getElementById('penyebab').value;
        const Korban = document.getElementById('korban').value;
        const tanggal = document.getElementById('tanggal').value;
        let waktuInput = document.getElementById('waktu').value;
        const lokasiPeruas = document.getElementById('lokasiPeruas').value;
        const idKendaraan = document.getElementById('combokendaran').value;
        const jenisJalur = document.querySelector('input[name="jenisJalur"]:checked').value;

        console.log("Selected idKendaraan: " + idKendaraan); // Log to confirm the selected value

        if (!idKendaraan) {
          alert("Please select a vehicle!");
          return;
        }

        let lokasiKecelakaan;
        if (document.getElementById('otherLocationCheckbox').checked) {
          lokasiKecelakaan = document.getElementById('otherLocation').value;
          if (!lokasiKecelakaan) {
            alert('Please provide the other location!');
            return;
          }
        } else {
          lokasiKecelakaan = document.getElementById('lokasikecelekaan').value;
          if (!lokasiKecelakaan) {
            alert('Please select a location!');
            return;
          }
        }

        // Ensure waktuInput is not empty and is in a valid format
        if (!waktuInput || !/^\d{2}:\d{2}$/.test(waktuInput)) {
          alert('Invalid time format. Please use HH:MM format.');
          return;
        }



        const formData = {
          penyebab,
          Korban,
          tanggal,
          waktuInput,
          idKendaraan: parseInt(idKendaraan, 10),
          lokasiKecelakaan,
          lokasiPeruas,
          jenisJalur
        };

        // Mengirim data ke backend menggunakan fetch
        fetch('/api/kecelakaan/', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData),
          })
          .then((response) => {
            if (!response.ok) {
              throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
          })
          .then((data) => {
            alert('Data submitted successfully!' + JSON.stringify(data));
            fetchAndLoadKecelakaan();
          })
          .catch((error) => {
            console.error('Error submitting form:', error);
            alert('An error occurred while submitting the form.');
          });
      });


    })
    .catch(function (error) {
      console.log('Error fetching data:', error);
    });

});

// Function to fetch and load accident data
function fetchAndLoadKecelakaan() {
  // Fetch accident data with pagination using fetch
  fetch(`/api/kecelakaan/?page=${currentPage}&limit=${limit}`)
    .then(response => {
      if (!response.ok) {
        throw new Error('Failed to fetch data');
      }
      return response.json();
    })
    .then(data => {
      console.log('Response from GetAllKecelakaan:', data);

      if (data && data.data) {
        let kecelakaanData = data.data;

        // Log data to check structure before sorting
        console.log('Before Sorting:', kecelakaanData.map(item => item.idKecelakaan));

        // Sort by accident ID (you can adjust this as needed)
        kecelakaanData.sort((a, b) => a.idKecelakaan - b.idKecelakaan);

        // Log sorted data
        console.log('After Sorting:', kecelakaanData.map(item => item.idKecelakaan));

        const tableData = kecelakaanData.map((kecelakaan, index) => ({
          'ID Kecelakaan': kecelakaan.idKecelakaan,
          'ID Kendaraan': kecelakaan.idKendaraan,
          'Penyebab': kecelakaan.penyebab,
          'Korban': kecelakaan.korban,
          'Tanggal': kecelakaan.tanggal,
          'Waktu': kecelakaan.waktu,
          'Lokasi Kecelakaan': kecelakaan.lokasiKecelakaan,
          'Lokasi Peruas': kecelakaan.lokasiPeruas,
          'Jenis Jalur': kecelakaan.jenisJalur,
          'Actions': index + 1
        }));

        // Load the table data into the UI
        loadTableD2(
          ['ID Kecelakaan', 'ID Kendaraan', 'Penyebab', 'Korban', 'Tanggal', 'Waktu', 'Lokasi Kecelakaan', 'Lokasi Peruas', 'Jenis Jalur' , 'Actions'],
          tableData
        );

        // Handle pagination
        handlePaginationKecelakaan(data.totalPages);
      } else {
        console.error('Invalid response data:', data);
      }
    })
    .catch(error => {
      console.error('Error fetching Kecelakaan data:', error);
      alert('Gagal memuat data Kecelakaan');
    });
}



function handlePaginationKecelakaan(totalPages) {
  const paginationContainer = document.getElementById('pagination-container');
  if (!paginationContainer) {
    console.error('Pagination container not found');
    return;
  }

  let paginationHTML = '';
  for (let i = 1; i <= totalPages; i++) {
    paginationHTML += `<button class="btn btn-sm btn-secondary" style="margin-right:10px;" onclick="goToPageKecelakaan(${i})">${i}</button>`;
  }

  paginationContainer.innerHTML = paginationHTML;
}

function goToPageKecelakaan(pageNumber) {
  currentPage = pageNumber;
  fetchAndLoadKecelakaan(); // Adjusted to call the correct function for loading data
}

// Function to load table data
function loadTableD2(headers, data) {
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
      // If the header is 'Actions', add the action buttons
      if (header === 'Actions') {
        tableHTML += `
                    <td>
                        <button class="btn btn-sm btn-danger" onclick="deleteRowKecelakaan('${row['ID Kecelakaan']}')">Delete</button>
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

function deleteRowKecelakaan(idKecelakaan) {
  if (confirm('Are you sure you want to delete this Kecelakaan?')) {
    fetch(`/api/kecelakaan/${idKecelakaan}`, {
        method: 'DELETE'
      })
      .then(response => {
        if (!response.ok) {
          throw new Error(`Failed to delete record with ID ${idKecelakaan}`);
        }
        return response.json();
      })
      .then(data => {
        if (data.message === "Kecelakaan deleted successfully") {
          alert('Kecelakaan deleted successfully');
          fetchAndLoadKecelakaan(); // Reload table data
        } else {
          alert('Failed to delete Kecelakaan');
        }
      })
      .catch(error => {
        console.error('Error deleting Kecelakaan:', error);
        alert('An error occurred while deleting the record.');
      });
  }
}


///////////////////////////////////////////////////////////////////////////// Kendaraan ///////////////////////////////////////////////////////////////////////////////////////

document.getElementById('form-jenis-kendaraan-tab').addEventListener('click', function () {
  document.getElementById('form-container1').style.display = 'none';
  document.getElementById('form-container3').style.display = 'none';

  // Show form-container2 and insert dynamic content into it
  document.getElementById('form-container2').style.display = 'block';
  document.getElementById('form-container2').innerHTML = `

     
  <form id="kendaraanForm">

  
  <div class="mb-3">
    <label for="namaKendaraana" class="form-label">Nama Kendaraan</label>
    <input type="text" class="form-control" id="namaKendaraana" placeholder="Enter Name">
  </div>
  <div class="mb-3">
    <label for="warnaa" class="form-label">Warna Kendaraan</label>
    <input type="text" class="form-control" id="warnaa" placeholder="Enter Color">
  </div>
  <div class="mb-3">
    <label for="tipea" class="form-label">Tipe Kendaraan</label>
   <input type="text" class="form-control" id="tipea" placeholder="Enter Type">
  </div>
  <div class="mb-3">
    <label for="platNomora" class="form-label">Plat Nomor</label>
   <input type="text" class="form-control" id="platNomora" placeholder="Enter Plate Number">
  </div>
  <button type="submit" class="btn btn-primary">Submit</button>
</form>

  `;
  fetchAndLoadKendaraan();

  console.log('Form innerHTML rendered');
  console.log(document.getElementById('namaKendaraan'));

  document.getElementById('kendaraanForm').addEventListener('submit', async function (e) {
    e.preventDefault(); // Prevent the default form submission behavior

    // Ambil nilai dari input form
    const namaKendaraan = document.getElementById('namaKendaraana').value;
    const warna = document.getElementById('warnaa').value;
    const tipe = document.getElementById('tipea').value;
    const platNomor = document.getElementById('platNomora').value;

    console.log('Form values:', {
      namaKendaraan,
      warna,
      tipe,
      platNomor,
    });



    // Kirim data ke endpoint API CreateKendaraan
    try {
      console.log('Sending request to /api/kendaraan/');

      const response = await fetch('/api/kendaraan/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          namaKendaraan: namaKendaraan,
          warna: warna,
          tipe: tipe,
          platNomor: platNomor,
        }),
      });

      console.log('Response status:', response.status);
      if (!response.ok) {
        const errorText = await response.text();
        console.log('Error response:', errorText);
        throw new Error(`Failed to create Kendaraan: ${errorText}`);
      }

      const result = await response.json();

      // Debug log untuk hasil respon dari server
      console.log('Server response:', result);

      alert('Kendaraan created successfully: ' + JSON.stringify(result));
      // Initial call to load the vehicle data
      fetchAndLoadKendaraan();

      // Refresh halaman setelah berhasil mendapatkan ID
      window.location.reload();
    } catch (error) {
      // Log error yang terjadi
      console.log('Error occurred:', error);
      alert('Error: ' + error.message);
    }
  });




});

async function fetchAndLoadKendaraan() {
  // Fetch the data for Kendaraan from the API with pagination
  try {
    const response = await fetch(`/api/kendaraan/?page=${currentPage}&limit=${limit}`);
    const data = await response.json();

    if (data && data.data) {
      let kendaraanData = data.data;

      // Optionally, you can sort the data based on a specific field
      kendaraanData.sort((a, b) => {
        // Sorting logic if needed, for example by name
        return a.namaKendaraan.localeCompare(b.namaKendaraan);
      });

      const tableData = kendaraanData.map((kendaraan, index) => ({
        'ID Kendaraan': kendaraan.idKendaraan,
        'Nama Kendaraan': kendaraan.namaKendaraan,
        'Warna': kendaraan.warna,
        'Tipe': kendaraan.tipe,
        'Plat Nomor': kendaraan.platNomor,
        'Actions': index + 1
      }));

      // Call the function to load the table
      loadKendaraanTable(
        ['ID Kendaraan', 'Nama Kendaraan', 'Warna', 'Tipe', 'Plat Nomor', 'Actions'],
        tableData
      );

      // Handle pagination if necessary
      handlePaginationKendaraan(data.totalPages);
    } else {
      console.error('Invalid response data:', data);
    }
  } catch (error) {
    console.error('Error fetching Kendaraan data:', error);
    alert('Gagal memuat data Kendaraan');
  }
}

// Pagination Handling
function handlePaginationKendaraan(totalPages) {
  const paginationContainer = document.getElementById('pagination-container');
  if (!paginationContainer) {
    console.error('Pagination container not found');
    return;
  }

  let paginationHTML = '';
  for (let i = 1; i <= totalPages; i++) {
    paginationHTML += `<button class="btn btn-sm btn-secondary" style="margin-right:10px;" onclick="goToPagekendaraan(${i})">${i}</button>`;
  }

  paginationContainer.innerHTML = paginationHTML;
}

function goToPagekendaraan(pageNumber) {
  currentPage = pageNumber;
  fetchAndLoadKendaraan();
}

function loadKendaraanTable(headers, data) {
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
      if (header === 'Actions') {
        tableHTML += `
          <td>
              <button class="btn btn-sm btn-primary" onclick="editKendaraanRow('${row['ID Kendaraan']}')">Edit</button>
              <button class="btn btn-sm btn-danger" onclick="deleteKendaraanRow('${row['ID Kendaraan']}')">Delete</button>
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

function editKendaraanRow(idKendaraan) {
  // Fetch the specific Kendaraan data
  fetch(`/api/kendaraan/${idKendaraan}`)
    .then(response => response.json())
    .then(data => {
      if (data) {
        // Populate a form or show an overlay with the vehicle data
        document.getElementById('namaKendaraan').value = data.namaKendaraan;
        document.getElementById('warna').value = data.warna;
        document.getElementById('tipe').value = data.tipe;
        document.getElementById('platNomor').value = data.platNomor;
        // Set the idKendaraan as a hidden field to keep it for submission
        document.getElementById('idKendaraan').value = data.idKendaraan;

        // Open the overlay for editing
        toggleOverlay(2);
      }
    })
    .catch(error => console.error('Error fetching Kendaraan data:', error));
}


// Function to update kendaraan
function updateKendaraan() {
  const idKendaraan = document.getElementById("idKendaraan").value; // Get the ID of the kendaraan
  const namaKendaraan = document.getElementById("namaKendaraan").value;
  const warna = document.getElementById("warna").value;
  const tipe = document.getElementById("tipe").value;
  const platNomor = document.getElementById("platNomor").value;

  // Validate that required fields are filled
  if (!idKendaraan || !namaKendaraan || !warna || !tipe || !platNomor) {
    alert("All fields are required!");
    return;
  }

  // Prepare the data to be sent
  const data = {
    namaKendaraan: namaKendaraan,
    warna: warna,
    tipe: tipe,
    platNomor: platNomor
  };

  // Send data to the server via an AJAX request (using fetch in this case)
  fetch(`/api/kendaraan/${idKendaraan}`, {
      method: 'PUT', // PUT request to update the data
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => {
      if (data.success) {
        alert('Kendaraan updated successfully!');
        // Optionally, redirect or refresh the page
        fetchAndLoadKendaraan();
      } else {
        alert('Failed to update kendaraan');
      }
    })
    .catch(error => {
      console.error('Error:', error);
      alert('Error updating kendaraan');
    });
}


function deleteKendaraanRow(idKendaraan) {
  if (confirm('Are you sure you want to delete this Kendaraan?')) {
    fetch(`/api/kendaraan/${idKendaraan}`, {
        method: 'DELETE'
      })
      .then(response => response.json())
      .then(data => {
        if (data.message === "Kendaraan soft-deleted successfully") {
          alert('Kendaraan deleted successfully');
          fetchAndLoadKendaraan(); // Reload the table after deletion
        } else {
          alert('Failed to delete Kendaraan');
        }
      })
      .catch(error => console.error('Error deleting Kendaraan:', error));
  }
}