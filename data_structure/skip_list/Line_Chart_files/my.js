var maxlevel = 32
var maxnum = 1024
var MONTHS = [];
var randData  = [];

for (i=0;i<maxnum;i++){
    MONTHS.push(i)
    randData.push(randme())
}

var config = {
    type: 'line',
    data: {
        labels: MONTHS,
        datasets: [{
            label: 'Demo',
            fill: false,
            backgroundColor: window.chartColors.blue,
            borderColor: window.chartColors.blue,
            data: randData
        }]
    },
    options: {
        responsive: true,
        title: {
            display: true,
            text: 'Skip list'
        },
        tooltips: {
            mode: 'index',
            intersect: false,
        },
        hover: {
            mode: 'nearest',
            intersect: true
        },
        scales: {
            xAxes: [{
                display: true,
                scaleLabel: {
                    display: true,
                    labelString: 'Month'
                }
            }],
            yAxes: [{
                display: true,
                scaleLabel: {
                    display: true,
                    labelString: 'Value'
                }
            }]
        }
    }
};

window.onload = function() {
    var ctx = document.getElementById('canvas').getContext('2d');
    window.myLine = new Chart(ctx, config);
};
function randme(){
    var level =0;
    while(Math.random() < 0.25){
        level++
    }
    return level < maxlevel  ? level: maxlevel
}