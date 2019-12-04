/**
* @file matchkey_test.go
* @brief Contains test functions for `matchkey.go`.
* @author Anadian
* @copyright MITlicensetm(2019,Anadian)
*/

// matchkey a datatype which is a superposition of a literal string, a wildcarded path, and/or a Regular Expression.
package matchkey;

//Dependencies
import(
	//Internal
	//Standard
	"testing"
	"log"
	//External
);

//Constants
const(
	//Exported Constants
	//Private Constants
);

//Types, structs, and methods
//Global Variables
var(
	//Exported Variables
	//Private Variables
);

//Exported Functions
func TestMatchKey(t *testing.T){
	log.SetFlags(LstdFlags|Lmicroseconds|Lshortfile);
	//Variables
	var matchkey_object [4]MatchKey_struct = nil;
	var new_error error_report.ErrorReport_struct;
	var match_bool bool = false;
	var match_error error = nil;
	//Expected returns
	var expected_matchkey_object [2]MatchKey_struct = { nil, (!nil) };
	var expected_new_error [3]error = { nil, invalid_param_error, compile_error;
	var expected_match_bool [2]bool = false;
	var expected_match_error [2]error = nil;
	//# Tests
	//## New
	//#### Params
	test_param [4]uint8 = {MATCHKEY_TYPE_STRING, MATCHKEY_TYPE_PATH;
	invalid_type_param uint8 = 4;
	//#### Invalid Argument
	//##### Conditions
	//###### Variables
	/*matchkey_object = nil;
	new_error = nil;
	match_bool = false;
	match_error = nil;
	//Expected returns
	expected_matchkey_object = nil;
	expected_new_error = nil;
	expected_match_bool bool = false;
	expected_match_error error = nil;*/
	/*
	for New
		if params
		then
			for match
				if params
				then
					success
				else params
					fail
		else params
			fail
			for match
				if params
				then
					success
				else
					fail
	*/
	//## New with an invalid matchkey_type
	matchkey_object, new_error = New( 4, "test*" );
	if( matchkey_object == nil ){
		if( new_error.CodeEqual( ERROR_CODE_INVALID_ARGUMENT_MATCHKEY_TYPE ){
		} else{
			log.Printf("Fail: new_error is nil.");
			t.Fail();
		}
	} else{
		log.Printf("Fail: matchkey_object isn't nil.");
		t.Fail();
	}
	//New: string-types
	matchkey_object, new_error = New( MATCHKEY_TYPE_STRING, "test*" );
	if( new_error == nil ){
		if( matchkey_object !== nil ){
			match_bool = matchkey_object.Match("test*");
			if( match_bool == true ){
				match_bool = matchkey_object.Match("test");
				if( match_bool == false ){
					match_bool = matchkey_object.Match("tests");
					if( match_bool == false ){
						log.Printf("Pass: string-type matching works.");
					} else{
						t.Fail();
						log.Printf("Fail");
					}
				} else{
					t.Fail();
					log.Printf("Fail");
				}
			} else{
				t.Fail();
				log.Printf("Fail");
			}
		} else{
			t.Fail();
			log.Printf("Fail");
		}
	} else{
		t.Fail();
		log.Printf("Fail");
	}
	//New: path-types
	matchkey_object, new_error = New( MATCHKEY_TYPE_PATH, "test*s+" );
	if( new_error == nil ){
		if( matchkey_object !== nil ){
			match_bool = matchkey_object.Match("test*");
			if( match_bool == false ){
				match_bool = matchkey_object.Match("test");
				if( match_bool == false ){
					match_bool = matchkey_object.Match("tests");
					if( match_bool == false ){
						match_bool = matchkey_object.Match("tes");
						if( match_bool == false ){
							match_bool = matchkey_object.Match("test*a");
							if( match_bool == false ){
								match_bool = matchkey_object.Match("testwords+");
								if( match_bool == true ){
									match_bool = matchkey_object.Match("testtttts+");
									if( match_bool == true ){
										match_bool
								log.Printf("Pass: path-type matching works.");
							} else{
								t.Fail();
								log.Printf("Fail");
							}
						} else{
							t.Fail();
							log.Printf("Fail");
						}
					} else{
						t.Fail();
						log.Printf("Fail");
					}
				} else{
					t.Fail();
					log.Printf("Fail");
				}
			} else{
				t.Fail();
				log.Printf("Fail");
			}
		} else{
			t.Fail();
			log.Printf("Fail");
		}
	} else{
		t.Fail();
		log.Printf("Fail");
	}
	//New: regex-types
	matchkey_object, new_error = New( MATCHKEY_TYPE_REGEX, "test*s+" );
	if( new_error == nil ){
		if( matchkey_object !== nil ){
			match_bool = matchkey_object.Match("test*");
			if( match_bool == true ){
				match_bool = matchkey_object.Match("test");
				if( match_bool == false ){
					match_bool = matchkey_object.Match("tests");
					if( match_bool == false ){
						log.Printf("Pass: regex-type matching works.");
					} else{
						t.Fail();
						log.Printf("Fail");
					}
				} else{
					t.Fail();
					log.Printf("Fail");
				}
			} else{
				t.Fail();
				log.Printf("Fail");
			}
		} else{
			t.Fail();
			log.Printf("Fail");
		}
	} else{
		t.Fail();
		log.Printf("Fail");
	}

//Private Functions

