let contactFormNode = document.getElementById("contactForm");
// console.log(contactFormNode);

contactFormNode.addEventListener("submit", (e) => {
  e.preventDefault();

  let formData = {
    nama: document.getElementById("nama").value.trim(),
    email: document.getElementById("email").value.trim(),
    phone: document.getElementById("phone").value.trim(),
    subject: document.getElementById("subject").value.trim(),
    pesan: document.getElementById("pesan").value.trim(),
  };

  let error = {};
  // Cek inputan, jika kosong maka error bernilai true
  for (data in formData) {
    formData[data] == "" ? (error[data] = true) : (error[data] = false);
  }

  // Cek semua error, dan ambil tindakan sesuai kondisi error tersebut
  error.nama || error.email || error.phone || error.subject || error.pesan
    ? alert("Semua Inputan Harus Diisi !")
    : sendMail(formData);
});

// Membuat fungsi untuk mengirim email
const sendMail = (data) => {
  // menyimpan dulu hasil data dari form ke dalam variabel, baru dipanggil saat membuka mailto
  let body = `Hello my name ${data.nama}, ${data.subject}, ${data.pesan}. I'm very happy if you want to contact me at ${data.phone} / ${data.email}.`;

  window.open(
    `mailto:sidikrudini16@gmail.com?subject=${data.subject}&body=${body}`
  );
};
