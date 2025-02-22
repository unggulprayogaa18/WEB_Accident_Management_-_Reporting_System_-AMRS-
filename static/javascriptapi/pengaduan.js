document.addEventListener('DOMContentLoaded', function () {

    // Show form-container3 and insert dynamic content into it
    document.getElementById('form-container').style.display = 'block';

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
    <label for="tanggal_waktu" class="form-label">Tanggal & Waktu Kejadian</label>
    <input type="datetime-local" class="form-control" id="tanggal_waktu">
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
 

  <input 
    type="text" 
    id="customKendaraanField" 
    class="form-control mt-2" 
    name="customKendaraan" 
    placeholder="format pengisian kendaraan : NAMA KENDARAAN - PLATNOMOR" 
    "
  >
</div>
 <div class="mb-3">
    <label for="jumlah_kendaraan" class="form-label">Jumlah Kendaraan yang Terlibat</label>
    <input type="number" class="form-control" id="jumlah_kendaraan" placeholder="Masukkan jumlah kendaraan">
  </div>


    
  <div class="mb-3">
    <label for="cuaca" class="form-label">Cuaca</label>
    <input type="text" class="form-control" id="cuaca" placeholder="Masukkan kondisi cuaca">
  </div>
  
  <div class="mb-3">
    <label for="jalur_tertutup" class="form-label">Jalur Tertutup Total</label>
    <select id="jalur_tertutup" class="form-select">
      <option value="ya">Ya</option>
      <option value="tidak">Tidak</option>
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
<label for="jenisJalurA">A (Bandung) </label>
</div>

<div>
<input type="radio" id="jenisJalurB" name="jenisJalur" value="B">
<label for="jenisJalurB">B (jakarta) </label>
</div>

<div>
<input type="radio" id="jenisJalurNone" name="jenisJalur" value="Bahu Jalur">
<label for="jenisJalurNone">Bahu Jalur</label>
</div>

</div>

<button type="submit" class="btn btn-primary">Submit</button>
</form>
`;

            // Insert the generated HTML into form-container3
            document.getElementById('form-container').innerHTML = formHTML;

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
            
                // Convert the datetime-local input to a full ISO string
                const tanggal_waktuInput = document.getElementById('tanggal_waktu').value;
                if (!tanggal_waktuInput) {
                    alert("Please enter the date and time of the incident.");
                    return;
                }
                const tanggal_waktu = new Date(tanggal_waktuInput).toISOString();
            
                // Get other form values and determine which fields to use
                const jumlah_kendaraan = document.getElementById('jumlah_kendaraan').value;
                const id_kendaraan = document.getElementById('jumlah_kendaraan').value;

                const cuaca = document.getElementById('cuaca').value;
                const jalurTertutup = document.getElementById('jalur_tertutup').value;
                const lokasiPeruas = document.getElementById('lokasiPeruas').value;
                const jenisJalur = document.querySelector('input[name="jenisJalur"]:checked') 
                                      ? document.querySelector('input[name="jenisJalur"]:checked').value 
                                      : null;
                const status_pengaduan = "belum_ditanggapi"; // Default status
            
                let lokasi_kecelakaan;
                if (document.getElementById('otherLocationCheckbox').checked) {
                    lokasi_kecelakaan = document.getElementById('otherLocation').value;
                    if (!lokasi_kecelakaan) {
                        alert('Please provide the other location!');
                        return;
                    }
                } else {
                    lokasi_kecelakaan = document.getElementById('lokasikecelekaan').value;
                    if (!lokasi_kecelakaan) {
                        alert('Please select a location!');
                        return;
                    }
                }
            
                // Create the form data object with proper keys
                const formData = {
                    tanggal_waktu,
                    lokasi_kecelakaan, // updated key
                    id_kendaraan,
                    jumlah_kendaraan,
                    cuaca,
                    jalur_tertutup_total: jalurTertutup,
                    lokasiPeruas,      // same as your struct tag ("lokasiPeruas")
                    jenisJalur,
                    status_pengaduan   // updated key
                };
            
                console.log("Form Data being sent:", formData);
            
                // Validate and parse jumlah_kendaraan as an integer
                const parsedJumlah = parseInt(formData.jumlah_kendaraan, 10);
                if (isNaN(parsedJumlah)) {
                    alert('Please enter a valid number for jumlah_kendaraan');
                    return;
                }
                formData.jumlah_kendaraan = parsedJumlah;
            
                // Send the data to the backend
                fetch('/api/Pengaduan/', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(formData)
                })
                .then(function (response) {
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    return response.json();
                })
                .then(function (responseData) {
                    console.log("Form submitted successfully:", responseData);
                    alert('BERHASIL MENYIMPAN DATA!');
                })
                .catch(function (error) {
                    console.error("Error submitting the form:", error);
                    alert('There was an error submitting the form.');
                });
            });
            
        });
    });