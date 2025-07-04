document.addEventListener('DOMContentLoaded', function () {
    const btn = document.querySelector('.btn');
    const moreText = document.querySelector('.more');

    btn.addEventListener('click', function (e) {
        e.preventDefault();

        moreText.classList.toggle('show');

        if (moreText.classList.contains('show')) {
            btn.textContent = 'کمتر بخوانید';

            btn.parentNode.appendChild(btn);
        } else {
            btn.textContent = 'بیشتر بخوانید';
        }
    });
});


