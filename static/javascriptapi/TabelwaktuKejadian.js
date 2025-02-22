var TabelwaktuKejadian = (function () {
    let currentPage = 1;
    const limit = 40;

    async function fetchDataAndRenderTable() {
        try {
            const response = await fetch(`/api/waktukejadian/?page=${currentPage}&limit=${limit}`);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            const results = data.data;
            const totalPages = data.totalPages;

            // Sort results based on waktu in ascending order
            results.sort((a, b) => {
                const timeA = new Date(`1970-01-01T${a.waktu}`); // Add a fake date to parse the time
                const timeB = new Date(`1970-01-01T${b.waktu}`); // Add a fake date to parse the time
                return timeA - timeB; // Sort by time
            });

            let tableHTML = '';
            results.forEach(item => {
                // Check if 'waktu' exists and split the time portion (HH:mm:ss)
                const formattedWaktu = item.waktu ? item.waktu.split(' ')[1] || 'N/A' : 'N/A'; // Extract time part (HH:mm:ss)

                // Add each row to the table
                tableHTML += `
                <tr>
    <td class="text-xs ps-4">${item.waktu}</td>
    <td class="text-xs ps-4">${item.totalKecelakaan || '0'} kecelakaan</td>
    <td class="text-xs ps-4">${item.namaKendaraan || 'N/A'}</td>
    <td class="text-xs ps-4">${item.tipeKendaraan || 'N/A'}</td>
    <td class="text-xs ps-4">${item.warnaKendaraan || 'N/A'}</td>
    <td class="text-xs ps-4">${item.platNomorKendaraan || 'N/A'}</td>
    <td class="text-xs ps-4">${item.namaLokasi || 'N/A'}</td> <!-- Added item -->
    <td class="text-xs ps-4">${item.jenisJalur || 'N/A'}</td> <!-- Added item -->
    <td class="text-xs ps-4">${item.penyebab || 'N/A'}</td> <!-- Added item -->
    <td class="text-xs ps-4">${item.korban || 'N/A'}</td> <!-- Added item -->
    <td class="text-xs ps-4">${item.lokasiKecelakaan || 'N/A'}</td> <!-- Added item -->

</tr>

            `;
            });

            // Render the table
            document.getElementById("table-body-ea").innerHTML = tableHTML;
            handlePagination(totalPages);

        } catch (error) {
            console.error("Error fetching data:", error);
            document.getElementById("table-body-ea").innerHTML = `<tr><td colspan="6" class="text-danger">Error fetching data. Please try again later.</td></tr>`;
        }
    }

    function handlePagination(totalPages) {
        const paginationContainer = document.getElementById("pagination-container");
        paginationContainer.innerHTML = '';

        for (let i = 1; i <= totalPages; i++) {
            const button = document.createElement('button');
            button.classList.add('btn', 'btn-sm', 'btn-secondary');
            button.style.marginRight = '10px';
            button.textContent = i;
            button.onclick = () => goToPage(i);
            paginationContainer.appendChild(button);
        }
    }

    function goToPage(pageNumber) {
        currentPage = pageNumber;
        fetchDataAndRenderTable();
    }

    // Initial data load
    fetchDataAndRenderTable();


    // Call the fetch function to load data when the page loads
    window.onload = function () {
        fetchDataAndRenderTable();
    }

    return {
        init: fetchDataAndRenderTable,
    };
})();