This project was for a very specific, narrow purpose; so the likelihood of this
coming in handy elsewhere is quite low.

A friend wanted some way to import the table at
https://die-deutsche-wirtschaft.de/das-ranking-der-groessten-mittelstaendler-deutschlands/
into a spreadsheet. But the table on the website is not only un-copyable, but
also un-selectable, making the task rather difficult. So he took screenshots of
the tables, saved them as pdf files and sent me the pdfs.

From the various pdf files, I extracted only pages containing the tables and
collated them into a single file called `combined.pdf`. The script here parses
that pdf file and outputs each row in the table as an individual record in the
`output.tsv` file.

The script isn't perfect, and sometimes fails to separate columns (but tells you
which rows these are, so you can manually fix them), or omits the last line at
the end of a page (which you need to check manually). It also doesn't write the
very last record in a separate line (also needs to be fixed manually).

After making all the necessary manual adjustments, run `check.go` to perform a
quick sanity check on the final TSV file.

The final TSV file can then be directly opened in your favourite spreadsheet
program.
