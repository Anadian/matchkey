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
	error_report "github.com/Anadian/error_report/source"
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
	log.SetFlags( log.LstdFlags | log.Lmicroseconds | log.Lshortfile);
	//Variables
	var matchkey_object MatchKey_struct;
	var matchkey_string_object MatchKey_struct;
	var matchkey_path_object MatchKey_struct;
	var matchkey_regex_object MatchKey_struct;
	var return_error error_report.ErrorReport_struct;
	//# Tests
	//## New
	//### Invalid matchkey_type
	matchkey_object, return_error = New( 4, "test*" );
	if( matchkey_object == MATCHKEY_NIL_VALUE ){
		if( return_error.CodeEqual( ERROR_CODE_INVALID_ARGUMENT_MATCHKEY_TYPE ) == true ){
			log.Printf("Success.");
		} else{
			log.Printf("Fail: return_error doesn't equal ERROR_CODE_INVALID_ARGUMENT_MATCHKEY_TYPE.");
			t.Fail();
		}
	} else{
		log.Printf("Fail: matchkey_object isn't nil.");
		t.Fail();
	}
	//### String Success
	matchkey_string_object, return_error = New( MATCHKEY_TYPE_STRING, "test*" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( matchkey_string_object != MATCHKEY_NIL_VALUE ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: matchkey_string_object is null.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: return_error isn't null.");
	}
	//### Path Failure
	/*matchkey_path_object, return_error = New( MATCHKEY_TYPE_PATH, "test*?- ]\\\\\\" ); //I can't find any pattern that will trigger the bad pattern error.
	if( return_error.CodeEqual( ERROR_CODE_PATH_MATCH ) == true ){
		if( matchkey_path_object == MATCHKEY_NIL_VALUE ){
			log.Printf("Success");
		} else{
			t.Fail();
			log.Printf("Fail: matchkey_path_object is null.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: return_error isn't ERROR_CODE_PATH_MATCH %v", return_error);
	}*/
	//### Path Success
	matchkey_path_object, return_error = New( MATCHKEY_TYPE_PATH, "test*" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( matchkey_path_object != MATCHKEY_NIL_VALUE ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: matchkey_path_object is null.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: return_error isn't null.");
	}
	//### Regex Failure
	matchkey_regex_object, return_error = New( MATCHKEY_TYPE_REGEX, "t^e(st*" ); //)
	if( matchkey_regex_object == MATCHKEY_NIL_VALUE ){
		if( return_error.CodeEqual( ERROR_CODE_REGEXP_COMPILE ) == true ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: return_error isn't ERROR_CODE_REGEXP_COMPILE.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: matchkey_regex_object isn't nil.");
	}
	//### Regex Success
	matchkey_regex_object, return_error = New( MATCHKEY_TYPE_REGEX, "test*" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( matchkey_regex_object != MATCHKEY_NIL_VALUE ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: matchkey_regex_object is null.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: return_error isn't null.");
	}
	//## Match
	//### Uninitialised
	var match_bool bool;
	match_bool, return_error = matchkey_object.Match( "test*" );
	if( return_error.CodeEqual( ERROR_CODE_INVALID_PROPERTY_MATCHKEY_TYPE ) == true ){
		log.Printf("Success.");
	} else{
		t.Fail();
		log.Printf("Fail: return_error wasn't ERROR_CODE_INVALID_PROPERTY_MATCHKEY_TYPE");
	}
	//### String Failure
	match_bool, return_error = matchkey_string_object.Match( "test" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == false ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: there was a match when there shouldn't have been one.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	match_bool, return_error = matchkey_string_object.Match( "testa" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == false ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: there was a match when there shouldn't have been one.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	//### String Success
	match_bool, return_error = matchkey_string_object.Match( "test*" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == true ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: there wasn't a match when there should have been one.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	//### Path Match Error
	/*match_bool, return_error = matchkey_path_object.Match( "test*?[" );
	if( return_error.CodeEqual( ERROR_CODE_PATH_MATCH ) == true ){
		log.Printf("Success.");
	} else{
		t.Fail();
		log.Printf("Fail: return_error wasn't ERROR_CODE_PATH_MATCH.");
	}*/
	//### Path Failure
	match_bool, return_error = matchkey_path_object.Match( "tes" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == false ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-positive match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	match_bool, return_error = matchkey_path_object.Match( "tesa" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == false ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-positive match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	//### Path Success
	match_bool, return_error = matchkey_path_object.Match( "test" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == true ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-negative match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	match_bool, return_error = matchkey_path_object.Match( "testaaaa" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == true ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-negative match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	//### Regex Failure
	match_bool, return_error = matchkey_regex_object.Match( "te" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == false ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-positive match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	//### Regex Success
	match_bool, return_error = matchkey_regex_object.Match( "tes" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == true ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-negative match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	match_bool, return_error = matchkey_regex_object.Match( "test" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == true ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-negative match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	match_bool, return_error = matchkey_regex_object.Match( "testttt" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == true ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-negative match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	match_bool, return_error = matchkey_regex_object.Match( "tesa" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == true ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-negative match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
	match_bool, return_error = matchkey_regex_object.Match( "testtttta" );
	if( return_error.CodeEqual( 0 ) == true ){
		if( match_bool == true ){
			log.Printf("Success.");
		} else{
			t.Fail();
			log.Printf("Fail: false-negative match.");
		}
	} else{
		t.Fail();
		log.Printf("Fail: there was an error.");
	}
}
//Private Functions

