const exportBtn = document.getElementById("export-btn");
const { width, height } = document
  .getElementById("cv-container")
  .getBoundingClientRect();

let filename = document.getElementById("employee").innerText;
filename = filename.replace(/\s+/, "_");
filename = filename.toLowerCase();
filename = filename + "_cv.pdf";

const opt = {
  margin: 0,
  filename: filename,
  image: { type: "jpeg", quality: 0.98 },
  html2canvas: { scale: 2, dpi: 300, letterRendering: true },
  jsPDF: { format: [width, height], unit: "px", orientation: "portrait" },
};

const container = document.getElementById("cv-container");

exportBtn.addEventListener("click", () => {
  html2pdf(container, opt);
});
