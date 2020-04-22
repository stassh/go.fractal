import * as React from "react";
import { ShoppingList } from "./components/ShopingList";
import { About } from "./components/About";
import { FractalGrid } from "./components/FractalGrid";
import { Rect } from "./model/Rect";


const AsyncImageLazy = React.lazy(() => import("./components/AsyncImage"));

export interface AppProps {}
export interface AppState {}

export class App extends React.Component<AppProps, AppState> {

    // private renderLoader = () => <div className="loader">loading</div>;
    private screenHeight = 0;
    private screenWidth = 0;
    render() {

        const src = "https://i.imgur.com/R5TraVR.png";
        const alt = "test";

        this.screenHeight = window.screen.height;
        this.screenWidth = window.screen.width;
        const clientRect = Rect.fromXY(this.screenWidth, this.screenHeight)

        return (
            <div>

                <About screenRect={clientRect} />

                <FractalGrid screenRect={clientRect}/>

                <h1>Width: {this.screenWidth} Height: {this.screenHeight} </h1>

                {/* <React.Suspense fallback={this.renderLoader()}>
                   <AsyncImageLazy src={src} alt={alt} />
                </React.Suspense> */}

            </div>
        );
    };
}