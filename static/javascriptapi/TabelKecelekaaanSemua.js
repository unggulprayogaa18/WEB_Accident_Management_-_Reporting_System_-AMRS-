var TabelKecelakaanSemua = (function () {
    function loadTableD2(columns, data) {
        const tableBody = document.getElementById('table-body2');
        tableBody.innerHTML = ''; // Clear previous rows

        // Loop through the data and add rows dynamically
        data.forEach(rowData => {
            const row = document.createElement('tr');

            // Create a cell for each column
            columns.forEach(column => {
                const cell = document.createElement('td');
                cell.classList.add('align-middle', 'text-left' , 'ps-4');
                cell.textContent = rowData[column] || ''; // Handle null values
                row.appendChild(cell);
            });

            // Add the row to the table
            tableBody.appendChild(row);
        });
    }

    function fetchAndLoadKecelakaan() {
        const currentPage = 1;
        const limit = 10;

        fetch(`/api/kecelakaan/semuatahun?page=${currentPage}&limit=${limit}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to fetch data');
                }
                return response.json();
            })
            .then(data => {
                console.log('Response from GetDatakecelakaanTahunini:', data);

                if (data && data.kecelakaan_data) { // Gunakan 'kecelakaan_data'
                    let kecelakaanData = data.kecelakaan_data;

                    // Group data by lokasiKecelakaan and month
                    let groupedData = {};

                    kecelakaanData.forEach(kecelakaan => {
                        const date = new Date(kecelakaan.tanggal);
                        const month = date.toLocaleString('default', {
                            month: 'long'
                        });

                        const key = `${kecelakaan.lokasiKecelakaan}_${month}`;

                        if (!groupedData[key]) {
                            groupedData[key] = {
                                lokasiKecelakaan: kecelakaan.lokasiKecelakaan,
                                tanggal: month,
                                jumlahKecelakaan: 0,
                            };
                        }

                        groupedData[key].jumlahKecelakaan += 1;
                    });

                    let tableData = Object.values(groupedData).map((kecelakaan, index) => ({
                        'Lokasi Kecelakaan': kecelakaan.lokasiKecelakaan,
                        'Bulan': kecelakaan.tanggal,
                        'Jumlah Kecelakaan': kecelakaan.jumlahKecelakaan,
                    }));

                    console.log('Formatted Kecelakaan Data:', tableData);

                    loadTableD2(
                        ['Lokasi Kecelakaan', 'Bulan', 'Jumlah Kecelakaan'],
                        tableData
                    );

                    // Handle pagination
                    handlePagination(data.total_pages); // Gunakan 'total_pages'
                } else {
                    console.error('Invalid response data:', data);
                }
            })
            .catch(error => {
                console.error('Error fetching Kecelakaan data:', error);
                alert('Gagal memuat data Kecelakaan');
            });
    }


    function handlePagination(totalPages) {
        const paginationContainer = document.getElementById('pagination-container2');
        if (!paginationContainer) {
            console.error('Pagination container not found');
            return;
        }

        let paginationHTML = '';
        for (let i = 1; i <= totalPages; i++) {
            paginationHTML += `<button class="btn btn-sm btn-secondary" style="margin-right:10px;" onclick="goToPage(${i})">${i}</button>`;
        }

        paginationContainer.innerHTML = paginationHTML;
    }

    function goToPage(pageNumber) {
        currentPage = pageNumber;
        fetchAndLoadLokasi();
    }

    // Call the fetch function to load data when the page loads
    window.onload = function () {
        fetchAndLoadKecelakaan();
    }
    return {
        init: fetchAndLoadKecelakaan,
    };
})();