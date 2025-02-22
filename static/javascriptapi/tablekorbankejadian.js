var tablekorbankejadian = (function () {
    let currentPage = 1;
    const limit = 40;

    async function fetchDataAndRenderTablekorbankejadian() {
        try {
            const response = await fetch(`/api/KejadianKorban/?page=${currentPage}&limit=${limit}`);
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
                // Check if 'korban' exists and split the string for selamat, meninggal, luka-luka
                let selamat = 0,
                    meninggal = 0,
                    lukaLuka = 0;
                if (item.korban) {
                    const korbanParts = item.korban.split(' - ');
                    korbanParts.forEach(part => {
                        const segments = part.trim().split(' ');
                        if (segments.length === 2) {
                            const count = parseInt(segments[0], 10);
                            const type = segments[1];
                            if (type === 'selamat') {
                                selamat += count;
                            } else if (type === 'meninggal') {
                                meninggal += count;
                            } else if (type === 'luka-luka') {
                                lukaLuka += count;
                            }
                        }
                    });
                }

                // Add each row to the table with calculated korban details
                tableHTML += `
                <tr>
                    <td class="text-xs ps-4">${item.waktu || 'N/A'}</td>
                    <td class="text-xs ps-4">${item.namaLokasi || 'N/A'}</td> <!-- Added item -->
                    <td class="text-xs ps-4">${item.jenisJalur || 'N/A'}</td> <!-- Added item -->
                    <td class="text-xs ps-4">${item.penyebab || 'N/A'}</td> <!-- Added item -->
                    <td class="text-xs ps-4">${item.korban || 'N/A'}</td> <!-- Added item -->
                    <td class="text-xs ps-4">${item.lokasiKecelakaan || 'N/A'}</td> <!-- Added item -->
                    <td class="text-xs ps-4">${selamat}</td> <!-- Selamat -->
                    <td class="text-xs ps-4">${meninggal}</td> <!-- Meninggal -->
                    <td class="text-xs ps-4">${lukaLuka}</td> <!-- Luka-luka -->
                </tr>
            `;
            });

            // Render the table
            document.getElementById("table-body-ea2").innerHTML = tableHTML;
            handlePagination(totalPages);

        } catch (error) {
            console.error("Error fetching data:", error);
            document.getElementById("table-body-ea2").innerHTML = `<tr><td colspan="6" class="text-danger">Error fetching data. Please try again later.</td></tr>`;
        }
    }

    function handlePagination(totalPages) {
        const paginationContainer = document.getElementById("pagination-container2");
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
        fetchDataAndRenderTablekorbankejadian();
    }

    // Initial data load
    fetchDataAndRenderTablekorbankejadian();

    // Call the fetch function to load data when the page loads
    window.onload = function () {
        fetchDataAndRenderTablekorbankejadian();
    }

    return {
        init: fetchDataAndRenderTablekorbankejadian,
    };
})();
