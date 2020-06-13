// chrome.runtime.sendMessage({greeting: 'hello'})

const prices = document.evaluate(
    '/html/body/div[1]/div[6]/div[1]/div[1]/div/div/div[3]/div/div/div/div[3]/div[4]/div[1]/div/div[1]/div/div/h3/span/span',
    document,
    null,
    XPathResult.ANY_TYPE,
    null,
);
console.log(prices);

const price = prices.iterateNext();
console.log(price);

if (price) {
    chrome.runtime.sendMessage({});
}

chrome.runtime.sendMessage({
    price: price.innerHTML,
});