let currentPage = 1;
const limit = 15;

async function fetchDataAndRenderTable() {
    try {
        const response = await fetch(`/api/Table/?page=${currentPage}&limit=${limit}`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        const results = data.data;
        const totalPages = data.totalPages;

        let tableHTML = '';
        results.forEach(item => {
            tableHTML += `
                <tr>
                    <td class="text-xs ps-4">${item.penyebab || 'N/A'}</td>
                    <td class="text-xs ps-4">${item.warna || 'N/A'}</td>
                    <td class="text-xs ps-4">${item.platNomor || 'N/A'}</td>
                    <td class="text-xs ps-4">${item.namaKendaraan || 'N/A'}</td>
                    <td class="text-xs ps-4">${item.korban || 'N/A'}</td>
                    <td class="text-xs ps-4">${new Date(item.tanggal).toLocaleDateString() || 'N/A'}</td>
                    <td class="text-xs ps-4">${new Date(item.waktu).toLocaleTimeString() || 'N/A'}</td>
                    <td class="text-xs ps-4">${item.namaLokasi || 'N/A'}</td>
                    <td class="text-xs ps-4">${item.lokasiKecelakaan || 'N/A'}</td>
                    <td class="text-xs ps-4">${item.jenisJalur || 'N/A'}</td>

                </tr>
            `;
        });

        // Render the table
        document.getElementById("table-body-ea").innerHTML = tableHTML;
        handlePagination(totalPages);

    } catch (error) {
        console.error("Error fetching data:", error);
        document.getElementById("table-body-ea").innerHTML = `<tr><td colspan="10" class="text-danger">Error fetching data. Please try again later.</td></tr>`;
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