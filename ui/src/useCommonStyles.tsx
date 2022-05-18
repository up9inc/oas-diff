import { makeStyles } from '@material-ui/core';

// @ts-ignore

export const useCommonStyles = makeStyles(() => ({
    button: {
        backgroundColor: "#205cf5",
        color: "white",
        fontWeight: "600 !important",
        fontSize: 12,
        padding: "7.5px 10px",
        borderRadius: "6px ! important",

        "&:hover": {
            backgroundColor: "#205cf5",
            cursor: "pointer"
        },
        "&:disabled": {
            backgroundColor: "rgba(0, 0, 0, 0.26)"
        }
    },
    select: {
        margin: "0px !important"
    }
}));
