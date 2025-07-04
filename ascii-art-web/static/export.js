let isExporting = false;

function initExport() {
  const exportBtn = document.getElementById('exportBtn');
  const toggleBtn = document.getElementById('exportToggleBtn');
  const optionsBox = document.getElementById('exportOptions');

  toggleBtn.addEventListener('click', () => {
    optionsBox.style.display = optionsBox.style.display === 'none' ? 'block' : 'none';
  });

  exportBtn.addEventListener('click', () => {
    if (isExporting) return;
    isExporting = true;
    setTimeout(() => isExporting = false, 1000);

    const output = document.querySelector('#asciiOutput .inner');
    if (!output || !output.textContent.trim()) {
      alert("Nothing to export.");
      return;
    }

    const format = document.getElementById('format').value;
    const filename = document.getElementById('filename').value.trim() || 'ascii-art-web-export';

    const formData = new FormData();
    formData.append("asciiText", output.textContent);
    formData.append("format", format);
    formData.append("filename", filename);

    fetch("/export", { method: "POST", body: formData })
      .then(res => {
        if (!res.ok) throw new Error("Export failed.");
        return res.blob();
      })
      .then(blob => {
        const a = document.createElement("a");
        a.href = URL.createObjectURL(blob);
        a.download = `${filename}.${format}`;
        a.click();
        URL.revokeObjectURL(a.href);
      })
      .catch(err => {
        alert("Export error: " + err.message);
      });
  });
}
