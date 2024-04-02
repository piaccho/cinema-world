import React, { useState } from "react";
import { Slide, Box, Stack, IconButton } from "@mui/material";
import { NavigateBefore, NavigateNext } from "@mui/icons-material";

interface CarouselProps {
    elements: React.ReactElement[];
}

const Carousel: React.FC<CarouselProps> = ({ elements }) => {
    const [currentPage, setCurrentPage] = useState(0);
    const [slideDirection, setSlideDirection] = useState<
        "left" | "right" | undefined
    >("right");

    const cardsPerPage = 6;

    const handleNextPage = () => {
        setSlideDirection("left");
        setCurrentPage((prevPage) => prevPage + 1);
    };

    const handlePrevPage = () => {
        setSlideDirection("right");
        setCurrentPage((prevPage) => prevPage - 1);
    };

    return (
        <Box
            sx={{
                display: "flex",
                flexDirection: "row",
                alignItems: "center",
                alignContent: "center",
                justifyContent: "center",
                height: "600px",
                
            }}
        >
            <IconButton
                onClick={handlePrevPage}
                disabled={currentPage === 0}
                sx={{ margin: 5 }}
            >
                <NavigateBefore />
            </IconButton>
            <Box
                sx={{
                    display: "flex",
                    flexDirection: "row",
                    alignItems: "center",
                    alignContent: "center",
                    justifyContent: "center",
                    height: "600px",
                }}
            >
                {elements.map((card, index) => (
                    <Box
                        key={`card-${index}`}
                        sx={{
                            width: "100%",
                            height: "100%",
                            display: currentPage === index ? "block" : "none",
                        }}
                    >
                        <Slide
                            direction={slideDirection}
                            in={currentPage === index}
                            timeout={1000}
                        >
                            <Stack
                                spacing={2}
                                direction="row"
                                alignItems="center"
                                justifyContent="center"
                            >
                                {elements.slice(
                                    index * cardsPerPage,
                                    index * cardsPerPage + cardsPerPage
                                )}
                            </Stack>
                        </Slide>
                    </Box>
                ))}
            </Box>
            <IconButton
                onClick={handleNextPage}
                disabled={currentPage >= Math.ceil(elements.length || 0) / cardsPerPage}
                sx={{ margin: 5 }}
            >
                <NavigateNext />
            </IconButton>
        </Box>
    );
};

export default Carousel;
