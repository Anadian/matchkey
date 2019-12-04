/**
* @file matchkey.go
* @brief A datatype which is a superposition of a literal string, a wildcarded path, and/or a Regular Expression.
* @author Anadian
* @copyright MITlicensetm(2019,Anadian)
*/

// matchkey a datatype which is a superposition of a literal string, a wildcarded path, and/or a Regular Expression.
package matchkey;

//# Dependencies
import(
	//## Internal
	//## Standard
	//"errors" //For catching Go's native errors.
	"fmt"
	//"strings"
	"path"
	"regexp"
	//## External
	error_report "github.com/Anadian/error_report/source"
);

//# Constants
const(
	//## Exported Constants
	//### Static Error Codes
	ERROR_CODE_INVALID_ARGUMENT_MATCHKEY_TYPE int64 = 2
	ERROR_CODE_REGEXP_COMPILE int64 = 16
	ERROR_CODE_INVALID_PROPERTY_MATCHKEY_TYPE int64 = 3
	ERROR_CODE_PATH_MATCH int64 = 17
	//### Non-Error Constants
	MATCHKEY_TYPE_UNSPECIFIED uint8 = 0 //Null
	MATCHKEY_TYPE_STRING uint8 = 1 //String literal
	MATCHKEY_TYPE_PATH uint8 = 2 //Path-expanded string
	MATCHKEY_TYPE_REGEX uint8 = 3 //Compiled Regular Expression
	/*MATCHKEY_TYPE_POSIX_REGEX uint8 3 //Compiled POSIX-Extended Regular Expression
	MATCHKEY_TYPE_RE2_REGEX uint8 = 4 //Compiled RE2-format Regular Expression*/
	//Private Constants
);

//Types, structs, and methods
// MatchKey_interface is the interface for `MatchKey`s
type MatchKey_interface interface{
	Match( name string ) (match_bool bool, return_error error)
	/* AsString() string //Returns the initial string value. */
}
// MatchKey_struct is the struct which implements the MatchKey interface and holds neccessary data for initialised MatchKeys.
type MatchKey_struct struct{
	Matchkey_type uint8
	Matchkey_string string
	compiled_regexp *regexp.Regexp
}

/**
* @fn Match
* @brief Returns true if the given string matches the MatchKey dependent on the MatchKey type; false otherwise.
* @struct matchkey MatchKey_struct
* @param check_string string [in] The string which will be matched against the MatchKey via the appropriate method for the MatchKey type.
* @return (bool, error_report.ErrorReport_struct)
* @retval true Match
* @retval false No match.
*/
// Match returns true if the given string matches the MatchKey dependent on the MatchKey type; false otherwise.
func (matchkey MatchKey_struct) Match( check_string string ) (match bool, return_error error_report.ErrorReport_struct){
	/* Variables */
	var path_match_error error = nil;
	/* Parametres */
	/* Function */
	switch(matchkey.Matchkey_type){
		case MATCHKEY_TYPE_STRING:
			if( check_string == matchkey.Matchkey_string ){
				match = true;
			} else{
				match = false;
			}
		case MATCHKEY_TYPE_PATH:
			match, path_match_error = path.Match( matchkey.Matchkey_string, check_string );
			if( path_match_error != nil ){
				return_error = error_report.New( ERROR_CODE_PATH_MATCH, map[string]interface{}{
					"message": fmt.Sprintf("Error: path.Match(\"%s\", \"%s\") returned an error: %s", matchkey.Matchkey_string, check_string, path_match_error),
					"matchkey_string": matchkey.Matchkey_string,
					"check_string": check_string,
					"path_match_error": path_match_error,
				}, nil );
			}
		case MATCHKEY_TYPE_REGEX:
			if( matchkey.compiled_regexp.MatchString(check_string) == true ){
				match = true;
			} else{
				match = false;
			}
		default:
			return_error = error_report.New( ERROR_CODE_INVALID_PROPERTY_MATCHKEY_TYPE, map[string]interface{}{
				"message": fmt.Sprintf("Error: invalid property `Matchkey_type` (%d); was this MatchKey properly initialised?", matchkey.Matchkey_type),
				"Matchkey_type": matchkey.Matchkey_type,
			}, nil );
	}
	/* Return */
	return;
}

//Global Variables
var(
	//Exported Variables
	MATCHKEY_NIL_VALUE MatchKey_struct = MatchKey_struct{ MATCHKEY_TYPE_UNSPECIFIED, "", nil }
	//Private Variables
);

//Exported Functions
/**
* @fn New
* @brief Creates a new MatchKey, for a specified type, from a given string value.
* @param matchkey_type uint8 [in] The type of the MatchKey: 1 for string-literal, 2 for wildcard-enabled path, and three for a full-on Regular Expression.
* @param matchkey_string string [in] The base string to be used for the MatchKey. If `matchkey_type` is `MATCHKEY_TYPE_REGEX` this string will be compiled into a RegExp object.
* @return (MatchKey_struct, error_report.ErrorReport_struct)
* @retval 0 Success
* @retval 1 Not Supported
* @retval >1 Error
*/
// New creates a new MatchKey, for a specified type, from a given string value.
func New( matchkey_type uint8, matchkey_string string ) (new_matchkey MatchKey_struct, return_error error_report.ErrorReport_struct){
	/* Variables */
	var match_bool bool;
	var match_error error = nil;
	var regexp_compile_error error = nil;
	/* Parametres */
	/* Function */
	switch(matchkey_type){
		case MATCHKEY_TYPE_STRING:
			new_matchkey.Matchkey_type = matchkey_type;
			new_matchkey.Matchkey_string = matchkey_string;
			new_matchkey.compiled_regexp = nil;
		case MATCHKEY_TYPE_PATH:
			match_bool, match_error = path.Match( matchkey_string, "" );
			if( match_error != nil ){
				return_error = error_report.New( ERROR_CODE_PATH_MATCH, map[string]interface{}{
					"match_error": match_error,
				}, nil );
				new_matchkey = MATCHKEY_NIL_VALUE;
			} else{
				return_error = error_report.New( 0, map[string]interface{}{
					"match_bool": match_bool,
					"match_error": match_error,
					"matchkey_string": matchkey_string,
				}, nil );
				new_matchkey.Matchkey_type = matchkey_type;
				new_matchkey.Matchkey_string = matchkey_string;
				new_matchkey.compiled_regexp = nil;
			}
		case MATCHKEY_TYPE_REGEX:
			new_matchkey.Matchkey_type = matchkey_type;
			new_matchkey.Matchkey_string = matchkey_string;
			new_matchkey.compiled_regexp, regexp_compile_error = regexp.Compile(matchkey_string);
			if( regexp_compile_error != nil ){
				return_error = error_report.New( ERROR_CODE_REGEXP_COMPILE, map[string]interface{}{
					"message": fmt.Sprintf("Error: regexp.Compile(%s) returned an error: %s", matchkey_string, regexp_compile_error.Error()),
					"matchkey_string": matchkey_string,
					"regexp_compile_error": regexp_compile_error,
				}, nil );
				new_matchkey = MATCHKEY_NIL_VALUE;
			}
		default:
			return_error = error_report.New( ERROR_CODE_INVALID_ARGUMENT_MATCHKEY_TYPE, map[string]interface{}{
				"message": fmt.Sprintf("Error: invalid argument `matchkey_type`: %d", matchkey_type),
				"matchkey_type": matchkey_type,
			}, nil );
			new_matchkey = MATCHKEY_NIL_VALUE;
	}
	/* Return */
	return;
}

//Private Functions

