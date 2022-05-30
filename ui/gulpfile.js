const gulp = require('gulp');
const inlinesource = require('gulp-inline-source');
const replace = require('gulp-replace');
const rename = require('gulp-rename');

const fileName = "template.html"
const fileLocation = '../static'

gulp.task('default', () => {
    return gulp
        .src('./build/*.html')
        .pipe(replace('.js"></script>', '.js" inline></script>'))
        .pipe(replace('rel="stylesheet">', 'rel="stylesheet" inline>'))
        .pipe(
            inlinesource({
                compress: false,
                ignore: ['png'],
            })
        )
        .pipe(rename(fileName))
        .pipe(gulp.dest(fileLocation));
});
