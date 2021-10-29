var React = require('react');
var Component = React.Component;
var CanvasJSReact = require('canvasjs-react-charts');
var CanvasJS = CanvasJSReact.CanvasJS;
var CanvasJSChart = CanvasJSReact.CanvasJSChart;

class Bars extends Component {

    constructor() {
        super();
        this.state = {
            dataPoints: [{ y: 0, label: "Kafka" },
            { y: 0, label: "RabbitMQ" },
            { y: 0, label: "Pub/Sub" }]
        };
    }

    componentDidMount() {
        let points = [];
        for (let i = 0; i < this.props.workers.length; i++) {
            const worker = this.props.workers[i];
            points.push({
                y: worker.count, label: worker._id
            });
        }
        if (points.length > 0) {
            this.setState({
                dataPoints: points
            });
        }
    }

    render() {
        const options = {
            animationEnabled: true,
            theme: "dark1",
            backgroundColor: "transparent",
            height: 365,
            title: {
                text: ""
            },
            axisX: {
                title: "",
                reversed: true
            },
            axisY: {
                title: "INSERTIONS",
                includeZero: true,
                labelFormatter: this.addSymbols
            },
            data: [{
                type: "bar",
                dataPoints: this.state.dataPoints
            }]
        }
        return (
            <CanvasJSChart options={options} />
        );
    }
    addSymbols(e) {
        var suffixes = ["", "K", "M", "B"];
        var order = Math.max(Math.floor(Math.log(e.value) / Math.log(1000)), 0);
        if (order > suffixes.length - 1)
            order = suffixes.length - 1;
        var suffix = suffixes[order];
        return CanvasJS.formatNumber(e.value / Math.pow(1000, order)) + suffix;
    }
}

export default Bars;