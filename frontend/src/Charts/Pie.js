var React = require('react');
var Component = React.Component;
var CanvasJSReact = require('canvasjs-react-charts');
var CanvasJSChart = CanvasJSReact.CanvasJSChart;

class Pie extends Component {

    constructor() {
        super();
        this.state = {
            dataPoints: []
        };
    }

    getTotal(top_3) {
        let total = 0;
        for (let i = 0; i < top_3.length; i++) {
            const count = top_3[i].count;
            total += count;
        }
        return total;
    }

    componentDidMount() {
        let points = [];
        let total = this.getTotal(this.props.top_3);
        for (let i = 0; i < this.props.top_3.length; i++) {
            const top = this.props.top_3[i];
            const y = (top.count / total) * 100;
            points.push({
                y: y.toFixed(2),
                label: top._id,
                indexLabelFontColor: "black"
            });
        }
        this.setState({
            dataPoints: points
        });
    }

    render() {
        const options = {
            theme: "dark1",
            backgroundColor: "#15181c",
            height: 325,
            animationEnabled: true,
            exportFileName: "Top 5 Hashtags",
            exportEnabled: true,
            title: {
                text: ""
            },
            legend: {
                fontColor: "white",
                padding: 8
            },
            data: [{
                type: "pie",
                showInLegend: true,
                legendText: "{label}",
                toolTipContent: "{label}: <strong>{y}%</strong>",
                indexLabel: "{y}%",
                indexLabelPlacement: "inside",
                dataPoints: this.state.dataPoints
            }]
        }
        return (
            <CanvasJSChart options={options} />
        );
    }
}

export default Pie;