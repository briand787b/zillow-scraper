import React from 'react';

import PropertyBoxHeader from './PropertyBoxHeader';
import PropertyBoxGeneral from './PropertyBoxGeneral';
import PropertyBoxPriceHistory from './PropertyBoxPriceHistory';
import PropertyBoxTopography from './PropertyBoxTopography';
import PropertyBoxConvenience from './PropertyBoxConvenience';
import '../styles/PropertyBox.css';

const detailTypes = {
    GENERAL: 'general',
    PRICE_HISTORY: 'price-history',
    TOPOGRAPHY: 'topography',
    CONVENIENCE: 'convenience',
}

class PropertyBox extends React.Component {
    state = {
        section: detailTypes.GENERAL,
    };

    handleChangeDetailType = (detailTypeStr) => {
        return () => {
            let detailType;
            for (const dtField in detailTypes) {
                if (dtField === detailTypeStr) {
                    detailType = detailTypes[dtField];
                    break;
                }
            }

            if (!detailType) {
                throw new Error(`incorrect detailTypeStr (${detailTypeStr}) passed to handleChangeDetailType`);
            }

            this.setState({ section: detailType });
        }
    }

    render() {
        let detailType;
        switch (this.state.section) {
            case detailTypes.PRICE_HISTORY:
                detailType = <PropertyBoxPriceHistory />;
                break;
            case detailTypes.GENERAL:
                detailType = <PropertyBoxGeneral />;
                break;
            case detailTypes.TOPOGRAPHY:
                detailType = <PropertyBoxTopography />;
                break;
            case detailTypes.CONVENIENCE:
                detailType = <PropertyBoxConvenience />;
                break;
            default:
                throw new Error(`invalid value for this.state.section: ${this.state.section}`);
        }

        return (
            <div className="property-box">
                <PropertyBoxHeader
                    detailTypes={detailTypes}
                    handleChangeDetailType={this.handleChangeDetailType}
                />
                <div className="property-box-detail">
                    {detailType}
                </div>
            </div>
        );
    }
}

export default PropertyBox;