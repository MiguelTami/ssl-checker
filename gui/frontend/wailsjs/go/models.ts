export namespace ssl {
	
	export class Endpoint {
	    ipAddress: string;
	    serverName: string;
	    statusMessage: string;
	    grade: string;
	
	    static createFrom(source: any = {}) {
	        return new Endpoint(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ipAddress = source["ipAddress"];
	        this.serverName = source["serverName"];
	        this.statusMessage = source["statusMessage"];
	        this.grade = source["grade"];
	    }
	}
	export class SSLResult {
	    host: string;
	    status: string;
	    endpoints: Endpoint[];
	
	    static createFrom(source: any = {}) {
	        return new SSLResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.host = source["host"];
	        this.status = source["status"];
	        this.endpoints = this.convertValues(source["endpoints"], Endpoint);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

