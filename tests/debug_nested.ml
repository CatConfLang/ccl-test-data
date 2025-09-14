open Ccl.Parser

let test_input = "database =\n  enabled = true\n  port = 5432"

let () = 
  Printf.printf "Input:\n%s\n\n" test_input;
  match parse test_input with
  | Ok entries ->
      Printf.printf "Parsed entries (%d):\n" (List.length entries);
      List.iteri (fun i entry ->
        Printf.printf "%d. key='%s' value='%s'\n" (i+1) entry.key entry.value
      ) entries
  | Error (`Parse_error msg) ->
      Printf.printf "Parse error: %s\n" msg
