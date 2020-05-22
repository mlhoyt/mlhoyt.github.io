#!/usr/bin/env perl

open(OFH, "> bookmarks_unprocessed.html") or do {
  printf("[FATAL] failed to open file for writing: bookmarks_unprocessed.html\n");
  exit(1);
};

while (<>) {
  if (/<A HREF=\"([^\"]+)\"\s+ADD_DATE=\"(\d+)\">([^<]+)<\/A>/) {
    my ($linkURL, $linkAddDate, $linkTitle) = (${1}, ${2}, ${3});
    printf("[DEBUG] LINK URL=${linkURL}\n");
    printf("[DEBUG] LINK ADD_DATE=${linkAddDate}\n");
    printf("[DEBUG] LINK TITLE=${linkTitle}\n");

    if (-d "./source/${linkAddDate}") {
      printf("[ERROR] directory already exists: ${linkAddDate}\n");
      continue;
    }

    `mkdir -p ./source/${linkAddDate}`;
    open(entryFH, "> ./source/${linkAddDate}/entry.md") or do {
      printf("[ERROR] failed to open file for writing: ./source/${linkAddDate}/entry.md\n");
      continue;
    };

    printf entryFH "---\n";
    printf entryFH "title: ${linkTitle}\n";
    printf entryFH "date: %s\n", epochToRFC(${linkAddDate});
    printf entryFH "tags:\n";
    printf entryFH "---\n";
    printf entryFH "\n";
    printf entryFH "[${linkTitle}](${linkURL})\n";
    close(entryFH);
  } else {
    printf OFH $_;
  }
}

close(OFH);

sub epochToRFC{
  my($epoch) = @_;

  my ($sec, $min, $hour, $day, $month, $year) = (gmtime($epoch))[0,1,2,3,4,5];
  return sprintf("%4d-%02d-%02dT%02d:%02d:%02d-00:00", ($year+1900), ($month+1), $day, $hour, $min, $sec)
}
