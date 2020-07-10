const url = location.href;
console.log('location.href: ', url);

const priceSpan = document.querySelector('.ds-value');
console.log('priceSpan', priceSpan);

const price = parseInt(priceSpan.innerHTML.slice(1).replace(',', ''));
console.log('price', price)




// const prices = document.evaluate(
//     '/html/body/div[1]/div[6]/div[1]/div[1]/div/div/div[3]/div/div/div/div[3]/div[4]/div[1]/div/div[1]/div/div/h3/span/span',
//     document,
//     null,
//     XPathResult.ANY_TYPE,
//     null,
// );
// console.log('prices: ', prices);

// const price = prices.iterateNext();
// console.log('price: ', price);

if (price) {
    chrome.runtime.sendMessage({
        price: price,
        url: url,
    });
} else {
    chrome.runtime.sendMessage({});
}